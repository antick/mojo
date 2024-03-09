const path = require('path');

const ck3GamePath = process.env.CK3_GAME_PATH;
const ck3GameModPath = path.join(process.env.CK3_GAME_USER_PATH, 'mod');

// Path of the folder where we are storing a copy of game files to track the changes that comes in new ck3 releases
const syncedCk3Folder = path.join('.build', 'game');

// Path of the folder where source code of the mods is stored
const modSourcePath = path.join('mods');

module.exports = {
  lastSyncedCk3Version: '1.12.2.1',
  localModBuildPath: '.build/mods',
  ck3GameModPath,
  syncedCk3Folder,
  modSourcePath,
  modDescriptorSourcePath: path.join(modSourcePath, 'descriptor.mod'),
  thumbnailSourcePath: path.join(modSourcePath, 'thumbnail.png'),
  modIdPrefix: 'antick',
  modFoldersToProcess: [
    'common',
    'events',
    'gfx',
    'gui',
    'history',
    'localization',
    'music'
  ],
  gameFoldersToSync: () => {
    const folders = [
      'common',
      'events',
      'gui',
      'history',
    ];

    let foldersToSync = {};

    folders.forEach((folder) => {
      foldersToSync[path.join(ck3GamePath, folder)] = path.join(syncedCk3Folder, folder);
    });

    foldersToSync[path.join(ck3GamePath, 'localization', 'english')] = path.join(syncedCk3Folder, 'localization', 'english');

    return foldersToSync;
  },
};
