const path = require('path');
const os = require('os');

function config() {
  const userHomeDir = os.homedir();

  // Path of the folder where CK3 is installed in your system
  const installedCk3GamePath = path.join('C:', 'Program Files (x86)', 'Steam', 'steamapps', 'common', 'Crusader Kings III', 'game');

  // Path of the folder where CK3 installs the custom user mods
  const gameCustomModPath = path.join(userHomeDir, 'OneDrive', 'Documents', 'Paradox Interactive', 'Crusader Kings III', 'mod');

  // Path of the folder where we are storing a copy of game files to track the changes that comes in new releases
  const modCk3Path = path.join('game', 'original-game-files');

  // Path of the folder where source code of the mods is stored
  const modPath = path.join('game', 'mojo-mods');

  return {
    SyncedCk3Version: '1.11.0.1',
    GameDataPath: path.join(userHomeDir, 'OneDrive', 'Documents', 'Paradox Interactive', 'Crusader Kings III'),
    ModBuildPathLocal: 'build/mods',
    ModPath: modPath,
    GameCustomModPath: gameCustomModPath,
    ModDescriptorSourcePath: path.join(modPath, 'descriptor.mod'),
    ThumbnailSourcePath: path.join(modPath, 'thumbnail.png'),
    OriginalCk3Path: installedCk3GamePath,
    ModCk3Path: modCk3Path,
    Ck3PullMapping: {
      [path.join(installedCk3GamePath, 'common')]: path.join(modCk3Path, 'common'),
      [path.join(installedCk3GamePath, 'events')]: path.join(modCk3Path, 'events'),
      [path.join(installedCk3GamePath, 'gui')]: path.join(modCk3Path, 'gui'),
      [path.join(installedCk3GamePath, 'history')]: path.join(modCk3Path, 'history'),
      [path.join(installedCk3GamePath, 'localization', 'english')]: path.join(modCk3Path, 'localization', 'english'),
    },
    ModIdPrefix: 'antick',
    ModFoldersToProcess: [
      'common',
      'events',
      'gfx',
      'gui',
      'history',
      'localization',
      'music'
    ]
  };
}

module.exports = config;
