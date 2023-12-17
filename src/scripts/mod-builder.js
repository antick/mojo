const fs = require('fs').promises;
const path = require('path');

// Configurations would be imported from your configuration files.
const config = require('../config').Config();
const modConfig = require('../config').ModConfig();

async function cleanup(modPath) {
  try {
    await fs.rm(modPath, { recursive: true, force: true });
    console.log(`🧹 Cleaned ${modPath} folder`);
  } catch (err) {
    console.error(`Error while removing ${modPath}:`, err);
    throw err;
  }
}

async function processDirectory(modBuildPath, srcDirectory, destDirectory, replacements) {
  await fs.mkdir(destDirectory, { recursive: true });

  const items = await fs.readdir(srcDirectory, { withFileTypes: true });

  for (const item of items) {
    const srcPath = path.join(srcDirectory, item.name);
    const destPath = path.join(destDirectory, item.name);

    if (item.isDirectory()) {
      await processDirectory(modBuildPath, srcPath, destPath, replacements);
    } else {
      await processFile(modBuildPath, srcPath, destPath, replacements);
    }
  }
}

async function processFile(modBuildPath, srcPath, destPath, replacements) {
  let destFile = path.basename(destPath);
  const formatToReplace = "_\\$%s\\$_"; // This may need to be adjusted for JavaScript regex syntax

  for (const [placeholder, replacement] of Object.entries(replacements)) {
    const re = new RegExp(formatToReplace.replace('%s', placeholder), 'g');
    destFile = destFile.replace(re, replacement);
  }

  const finalDestPath = path.join(path.dirname(destPath), destFile);

  try {
    if (await fs.stat(finalDestPath)) {
      await fs.rm(finalDestPath);
    }
  } catch (err) {
    // File does not exist, ignore
  }

  // Skip text replacement for certain file types
  const ext = path.extname(srcPath);
  if (ext === ".dds" || ext === ".png") {
    return copyFiles(srcPath, finalDestPath);
  }

  try {
    let data = await fs.readFile(srcPath, 'utf8');

    for (const [placeholder, replacement] of Object.entries(replacements)) {
      const re = new RegExp(formatToReplace.replace('%s', placeholder), 'g');
      data = data.replace(re, replacement);
    }

    await fs.mkdir(modBuildPath, { recursive: true });
    await fs.writeFile(finalDestPath, data);
  } catch (err) {
    console.error("Error processing file:", err);
    throw err;
  }
}

const copyFiles = async (src, dst) => {
  try {
    const data = await fs.readFile(src);
    await fs.writeFile(dst, data);
  } catch (err) {
    console.error("Error while copying the file:", err);
    throw err;
  }
};

const build = async (modBuildPath, modName, replacements) => {
  const filesAndFolderMapping = {};

  // Assuming 'config.ModFoldersToProcess' and 'config.ModPath' are defined
  config.ModFoldersToProcess.forEach(value => {
    filesAndFolderMapping[path.join(config.ModPath, modName, value)] = path.join(modBuildPath, value);
  });

  const textReplacement = mergeReplacements(replacements, baseReplacements());

  for (const [srcPath, destPath] of Object.entries(filesAndFolderMapping)) {
    let info;
    try {
      info = await fs.stat(srcPath);
    } catch (err) {
      if (err.code === 'ENOENT') {
        // File or directory does not exist, skip
        continue;
      } else {
        console.error("Build failed while accessing the given path:", err);
        throw err;
      }
    }

    if (info.isDirectory()) {
      try {
        processDirectory(modBuildPath, srcPath, destPath, textReplacement); // Function not provided, assuming it exists
      } catch (err) {
        console.error("Build failed for the given directory:", err);
        throw err;
      }
    } else {
      try {
        processFile(modBuildPath, srcPath, destPath, textReplacement);
      } catch (err) {
        console.error("Build failed:", err);
        throw err;
      }
    }
  }
};

function mergeReplacements(map1, map2) {
  const mergedMap = {};
  for (const [key, value] of Object.entries(map1)) {
    mergedMap[key] = value;
  }
  for (const [key, value] of Object.entries(map2)) {
    mergedMap[key] = value;
  }
  return mergedMap;
}

function baseReplacements() {
  return {};
}

async function buildCombinedDescriptorFile(modBuildPath) {
  try {
    const err = await processFile(
      modBuildPath,
      config.ModDescriptorSourcePath,
      path.join(modBuildPath, 'descriptor.mod'),
      modConfig.CombinedMod.Replacements
    );

    return err;
  } catch (error) {
    console.error('Error in buildCombinedDescriptorFile:', error);
    throw error;
  }
}

async function buildLooseDescriptorFiles(modBuildPath, replacements) {
  try {
    const err = await processFile(
      modBuildPath,
      config.ModDescriptorSourcePath,
      path.join(modBuildPath, 'descriptor.mod'),
      replacements
    );

    return err;
  } catch (error) {
    console.error('Error in buildLooseDescriptorFiles:', error);
    throw error;
  }
}

async function buildModFile(modBuildPath, modFileName) {
  try {
    const sourceModFilePath = path.join(config.ModPath, modConfig.CombinedMod.Replacements['modFolderName'] + '.mod');
    const err = await processFile(
      modBuildPath,
      sourceModFilePath,
      path.join(modBuildPath, modFileName),
      modConfig.CombinedMod.Replacements
    );

    return err;
  } catch (error) {
    console.error('Error in buildModFile:', error);
    throw error;
  }
}

async function buildLooseModFiles(modBuildPath, modFileName, replacements) {
  try {
    const sourceModFilePath = path.join(config.ModPath, modConfig.CombinedMod.Replacements['modFolderName'] + '.mod');
    const err = await processFile(
      modBuildPath,
      sourceModFilePath,
      path.join(modBuildPath, modFileName),
      replacements
    );

    return err;
  } catch (error) {
    console.error('Error in buildLooseModFiles:', error);
    throw error;
  }
}

async function buildCombinedThumbnailFile(modBuildPath) {
  try {
    const err = await processFile(
      modBuildPath,
      config.ThumbnailSourcePath,
      path.join(modBuildPath, 'thumbnail.png'),
      {}
    );

    return err;
  } catch (error) {
    console.error('Error in buildCombinedThumbnailFile:', error);
    throw error;
  }
}

async function buildLooseThumbnailFiles(modBuildPath, modFolderName) {
  try {
    const err = await processFile(
      modBuildPath,
      path.join(config.ModPath, modFolderName, 'thumbnail.png'),
      path.join(modBuildPath, 'thumbnail.png'),
      {}
    );

    return err;
  } catch (error) {
    console.error('Error in buildLooseThumbnailFiles:', error);
    throw error;
  }
}

async function buildCombinedMod(modBuildPath, selectedModKeys) {
  console.log("-----------------------------------");
  try {
    await cleanup(modBuildPath);
    await buildCombinedDescriptorFile(modBuildPath);
    await buildCombinedThumbnailFile(modBuildPath);

    for (const modKey of selectedModKeys) {
      const modDetails = modConfig.SubMods[modKey];
      if (!modDetails) {
        console.log(`Mod ${modKey} not found`);
        continue;
      }

      if (!modDetails.Enabled) {
        console.log(`❗️${modKey} is disabled, skipping`);
        continue;
      } else {
        console.log(`📦 Building ${modKey}`);
      }

      await build(modBuildPath, modKey, modDetails.Replacements);
    }

    console.log("-----------------------------------");
    console.log(`✅ Build successful in ${modBuildPath} folder`);
    console.log("-----------------------------------");
  } catch (err) {
    throw err;
  }
}

async function buildLooseMods(buildPath, selectedModKeys) {
  console.log("-----------------------------------");
  try {
    for (const modKey of selectedModKeys) {
      const modDetails = modConfig.SubMods[modKey];
      if (!modDetails) {
        console.log(`Mod ${modKey} not found`);
        continue;
      }

      if (!modDetails.Enabled) {
        console.log(`❗️${modKey} is disabled, skipping`);
        continue;
      } else {
        console.log(`📦 Building ${modKey}`);
      }

      const modFolderName = modDetails.Replacements["modFolderName"];
      const modBuildPath = path.join(buildPath, modFolderName);
      const modFileName = modFolderName + ".mod";

      console.log("-----------------------------------");
      await cleanup(modBuildPath);
      await buildLooseModFiles(buildPath, modFileName, modDetails.Replacements);
      await buildLooseDescriptorFiles(modBuildPath, modDetails.Replacements);
      await buildLooseThumbnailFiles(modBuildPath, modFolderName);
      await build(modBuildPath, modKey, modDetails.Replacements);

      console.log(`✅ Build successful in ${modBuildPath} folder`);
    }
    console.log("-----------------------------------");
  } catch (err) {
    throw err;
  }
}

module.exports = {
  build,
  cleanup,
  processFile,
  processDirectory,
  buildCombinedDescriptorFile,
  buildLooseDescriptorFiles,
  buildModFile,
  buildLooseModFiles,
  buildCombinedThumbnailFile,
  buildLooseThumbnailFiles
};
