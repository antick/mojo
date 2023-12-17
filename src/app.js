const path = require('path');
const config = require('./config');
const scripts = require('./scripts');

const modConfig = config.ModConfig();

async function BuildModsInLocal(selectedModKeys, buildCombinedMod) {
  if (!selectedModKeys.length) {
    throw new Error('No mods selected for building');
  }

  if (buildCombinedMod) {
    const modFileName = `${modConfig.CombinedMod.Replacements.modFolderName}.mod`;

    try {
      await scripts.BuildModFile(config.ModBuildPathLocal, modFileName);
    } catch (error) {
      console.error('Error while processing mojo.mod file:', error);
      return error;
    }

    const modBuildPath = path.join(config.ModBuildPathLocal, modConfig.CombinedMod.Replacements.modFolderName);

    try {
      await scripts.BuildCombinedMod(modBuildPath, selectedModKeys);
    } catch (error) {
      console.error(`Error building combined mod in local: ${error}`);
      return error;
    }
  } else {
    try {
      await scripts.BuildLooseMods(config.ModBuildPathLocal, selectedModKeys);
    } catch (error) {
      console.error(`Error building selected mod(s) in local: ${error}`);
      return error;
    }
  }

  return null;
}

async function BuildModsInGame(selectedModKeys, buildCombinedMod) {
  if (!selectedModKeys.length) {
    throw new Error('No mods selected for building');
  }

  if (buildCombinedMod) {
    const modFileName = `${modConfig.CombinedMod.Replacements.modFolderName}.mod`;

    try {
      await scripts.BuildModFile(config.GameCustomModPath, modFileName);
    } catch (error) {
      console.error('Error while processing mojo.mod file:', error);
      return error;
    }

    const modBuildPath = path.join(config.GameCustomModPath, modConfig.CombinedMod.Replacements.modFolderName);

    try {
      await scripts.BuildCombinedMod(modBuildPath, selectedModKeys);
    } catch (error) {
      console.error(`Error building combined mod in game: ${error}`);
      return error;
    }
  } else {
    try {
      await scripts.BuildLooseMods(config.GameCustomModPath, selectedModKeys);
    } catch (error) {
      console.error(`Error building selected mod(s) in game: ${error}`);
      return error;
    }
  }

  return null;
}

async function PullCk3GameFiles() {
  try {
    await scripts.Pull(config.Ck3PullMapping);
    console.log(`📦 Pulled CK3 game files to ${configuration.ModCk3Path}`);
  } catch (error) {
    console.error(`Error pulling CK3 game files: ${error}`);
    return error;
  }

  console.log(` Pulled CK3 game files to ${config.ModCk3Path}`);
  return null;
}

async function GetModList() {
  return modConfig;
}

module.exports = {
  BuildModsInLocal,
  BuildModsInGame,
  PullCk3GameFiles,
  GetModList,
};
