const fs = require('fs').promises;
const path = require('path');
const config = require('../config');
const modConfig = require('../config/mods');

async function cleanup(modPath) {
  try {
    await fs.rm(modPath, {recursive: true, force: true});
    console.log(`🧹 Cleaned ${modPath} folder`);
  } catch (err) {
    console.error(`Error while removing ${modPath}:`, err);
    throw err;
  }
}

async function processDirectory(modBuildPath, srcDirectory, destDirectory, replacements) {
  await fs.mkdir(destDirectory, {recursive: true});

  const items = await fs.readdir(srcDirectory, {withFileTypes: true});

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
  const formatToReplace = '_\\$%s\\$_';

  for (const [placeholder, replacement] of Object.entries(replacements)) {
    const re = new RegExp(formatToReplace.replaceAll('%s', placeholder), 'g');
    destFile = destFile.replaceAll(re, replacement);
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
  if (ext === '.dds' || ext === '.png') {
    return copyFiles(srcPath, finalDestPath);
  }

  try {
    let data = await fs.readFile(srcPath, 'utf8');

    for (const [placeholder, replacement] of Object.entries(replacements)) {
      const re = new RegExp(formatToReplace.replaceAll('%s', placeholder), 'g');
      data = data.replaceAll(re, replacement);
    }

    await fs.mkdir(modBuildPath, {recursive: true});
    await fs.writeFile(finalDestPath, data);
  } catch (err) {
    console.error('Error processing file:', err);
    throw err;
  }
}

const copyFiles = async (src, dst) => {
  try {
    const data = await fs.readFile(src);
    await fs.writeFile(dst, data);
  } catch (err) {
    console.error('Error while copying the file:', err);
    throw err;
  }
};

const build = async (modBuildPath, modName, replacements) => {
  const filesAndFolderMapping = {};

  config.modFoldersToProcess.forEach(value => {
    filesAndFolderMapping[path.join(config.modSourcePath, modName, value)] = path.join(modBuildPath, value);
  });

  const textReplacement = mergeReplacements(replacements, {});

  for (const [srcPath, destPath] of Object.entries(filesAndFolderMapping)) {
    let info;
    try {
      info = await fs.stat(srcPath);
    } catch (err) {
      if (err.code === 'ENOENT') {
        // File or directory does not exist, skip
        continue;
      } else {
        console.error('Build failed while accessing the given path:', err);
        throw err;
      }
    }

    if (info.isDirectory()) {
      try {
        await processDirectory(modBuildPath, srcPath, destPath, textReplacement);
      } catch (err) {
        console.error('Build failed for the given directory:', err);
        throw err;
      }
    } else {
      try {
        await processFile(modBuildPath, srcPath, destPath, textReplacement);
      } catch (err) {
        console.error('Build failed:', err);
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

async function buildCombinedDescriptorFile(modBuildPath) {
  try {
    return await processFile(
      modBuildPath,
      config.modDescriptorSourcePath,
      path.join(modBuildPath, 'descriptor.mod'),
      modConfig.combinedMod
    );
  } catch (error) {
    console.error('Error in buildCombinedDescriptorFile:', error);
    throw error;
  }
}

async function buildLooseDescriptorFiles(modBuildPath, replacements) {
  try {
    return await processFile(
      modBuildPath,
      config.modDescriptorSourcePath,
      path.join(modBuildPath, 'descriptor.mod'),
      replacements
    );
  } catch (error) {
    console.error('Error in buildLooseDescriptorFiles:', error);
    throw error;
  }
}

async function buildModFile(modBuildPath, modFileName) {
  try {
    const sourceModFilePath = path.join(config.modSourcePath, modConfig.combinedMod['modFolderName'] + '.mod');
    return await processFile(
      modBuildPath,
      sourceModFilePath,
      path.join(modBuildPath, modFileName),
      modConfig.combinedMod
    );
  } catch (error) {
    console.error('Error in buildModFile:', error);
    throw error;
  }
}

async function buildLooseModFiles(modBuildPath, modFileName, replacements) {
  try {
    const sourceModFilePath = path.join(config.modSourcePath, modConfig.combinedMod['modFolderName'] + '.mod');

    return await processFile(
      modBuildPath,
      sourceModFilePath,
      path.join(modBuildPath, modFileName),
      replacements
    );
  } catch (error) {
    console.error('Error in buildLooseModFiles:', error);
    throw error;
  }
}

async function buildCombinedThumbnailFile(modBuildPath) {
  try {
    return await processFile(
      modBuildPath,
      config.thumbnailSourcePath,
      path.join(modBuildPath, 'thumbnail.png'),
      {}
    );
  } catch (error) {
    console.error('Error in buildCombinedThumbnailFile:', error);
    throw error;
  }
}

async function buildLooseThumbnailFiles(modBuildPath, modFolderName) {
  try {
    return await processFile(
      modBuildPath,
      path.join(config.modSourcePath, modFolderName, 'thumbnail.png'),
      path.join(modBuildPath, 'thumbnail.png'),
      {}
    );
  } catch (error) {
    console.error('Error in buildLooseThumbnailFiles:', error);
    throw error;
  }
}

async function buildCombinedMod(modBuildPath, selectedModKeys) {
  console.log('-----------------------------------');
  try {
    await cleanup(modBuildPath);
    await buildCombinedDescriptorFile(modBuildPath);
    await buildCombinedThumbnailFile(modBuildPath);

    const subMods = !selectedModKeys.length ? Object.keys(modConfig.subMods) : selectedModKeys;

    for (const modKey of subMods) {
      const modDetails = modConfig.subMods[modKey];
      if (!modDetails) {
        console.log(`Mod ${modKey} not found`);
        continue;
      }

      console.log(`📦 Building ${modKey}`);

      await build(modBuildPath, modKey, modDetails);
    }

    console.log('-----------------------------------');
    console.log(`✅ Build successful in ${modBuildPath} folder`);
    console.log('-----------------------------------');
  } catch (err) {
    throw err;
  }
}

async function buildLooseMods(buildPath, selectedModKeys) {
  console.log('-----------------------------------');
  try {
    const subMods = !selectedModKeys.length ? Object.keys(modConfig.subMods) : selectedModKeys;

    for (const modKey of subMods) {
      const modDetails = modConfig.subMods[modKey];
      if (!modDetails) {
        console.log(`Mod ${modKey} not found`);
        continue;
      }

      console.log(`📦 Building ${modKey}`);

      const modFolderName = modDetails['modFolderName'];
      const modBuildPath = path.join(buildPath, modFolderName);
      const modFileName = modFolderName + '.mod';

      console.log('-----------------------------------');
      await cleanup(modBuildPath);
      await buildLooseModFiles(buildPath, modFileName, modDetails);
      await buildLooseDescriptorFiles(modBuildPath, modDetails);
      await buildLooseThumbnailFiles(modBuildPath, modFolderName);
      await build(modBuildPath, modKey, modDetails);

      console.log(`✅ Build successful in ${modBuildPath} folder`);
    }
    console.log('-----------------------------------');
  } catch (err) {
    throw err;
  }
}

module.exports = {
  build,
  cleanup,
  buildCombinedMod,
  buildModFile,
  buildLooseMods,
};
