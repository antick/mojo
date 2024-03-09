const path = require('path');
const config = require('./');

const combinedModFolderName = 'mojo';
const modBuildPath = path.join(config.ck3GameModPath, combinedModFolderName);

module.exports = {
  combinedMod: {
    modFolderName: combinedModFolderName,
    modName: 'Mojo by Antick',
    modVersion: '0.1.0',
    supportedGameVersion: '1.11.0.1',
    modRemoteFileId: '', // 'remote_file_id="000000000"'
    modBuildPath: modBuildPath.replaceAll(/\\/g, '/'),
    modTags: `{
	"Alternative History"
	"Character Interactions"
	"Culture"
	"Decisions"
	"Events"
	"Fixes"
	"Gameplay"
	"Historical"
	"Utilities"
	"Warfare"
}`,
  },

  subMods: {
    'age-of-invasion': {
      modId: `${config.modIdPrefix}_age_of_invasion`,
      modFolderName: 'age-of-invasion',
      modName: 'Age of Invasion',
      modVersion: '1.0.0',
      supportedGameVersion: '1.12.2.1',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'age-of-invasion').replaceAll(/\\/g, '/'),
      modTags: `{
	"Gameplay"
	"Utilities"
	"Warfare"
}`,
    },

    'grand-council': {
      modId: `${config.modIdPrefix}_grand_council`,
      modFolderName: 'grand-council',
      modName: 'Grand Council',
      modVersion: '1.0.0',
      supportedGameVersion: '1.10.0.1',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'grand-council').replaceAll(/\\/g, '/'),
      modTags: `{
	"Gameplay"
	"Utilities"
}`,
    },

    'refurbished-titles': {
      modId: `${config.modIdPrefix}_refurbished_titles`,
      modFolderName: 'refurbished-titles',
      modName: 'Refurbished Titles',
      modVersion: '1.0.0',
      supportedGameVersion: '1.10.0.1',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'refurbished-titles').replaceAll(/\\/g, '/'),
      modTags: `{
	"Culture"
	"Fixes"
}`,
    },

    'tweak-n-treat': {
      modId: `${config.modIdPrefix}_tweak_n_treat`,
      modFolderName: 'tweak-n-treat',
      modName: 'Tweak n Treat',
      modVersion: '1.0.0',
      supportedGameVersion: '1.10.0.1',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'tweak-n-treat').replaceAll(/\\/g, '/'),
      modTags: `{
	"Fixes"
}`,
    },

    'united-thrones': {
      modId: `${config.modIdPrefix}_united_thrones`,
      modFolderName: 'united-thrones',
      modName: 'United Thrones',
      modVersion: '1.0.0',
      supportedGameVersion: '1.10.0.1',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'united-thrones').replaceAll(/\\/g, '/'),
      modTags: `{
	"Fixes"
}`,
    },

    'let-there-be-music': {
      modId: `${config.modIdPrefix}_let_there_be_music`,
      modFolderName: 'let-there-be-music',
      modName: 'Let there be music',
      modVersion: '1.0.0',
      supportedGameVersion: '1.10.0.1',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'let-there-be-music').replaceAll(/\\/g, '/'),
      modTags: `{
	"Utilities"
}`,
    },

    'auto-pause-game': {
      modId: `${config.modIdPrefix}_auto_pause`,
      modFolderName: 'auto-pause-game',
      modName: 'Auto Pause Game',
      modVersion: '1.0.1',
      supportedGameVersion: '1.11.0.1',
      modRemoteFileId: 'remote_file_id="2906586207"',
      modBuildPath: path.join(config.ck3GameModPath, 'auto-pause-game').replaceAll(/\\/g, '/'),
      modTags: `{
	"Fixes"
	"Utilities"
}`,
    },

    'epe-barbershop-idle-pose-compatch': {
      modId: `${config.modIdPrefix}_idle_pose_compatch`,
      modFolderName: 'epe-barbershop-idle-pose-compatch',
      modName: 'EPE + Barbershop Idle Pose Compatch By Antick',
      modVersion: '1.1.1',
      supportedGameVersion: '1.8.1',
      modRemoteFileId: 'remote_file_id="2870777363"',
      modBuildPath: path.join(config.ck3GameModPath, 'epe-barbershop-idle-pose-compatch').replaceAll(/\\/g, '/'),
      modTags: `{
	"Fixes"
	"Portraits"
}`,
    },

    'traits-mania': {
      modId: `${config.modIdPrefix}_traits_mania`,
      modFolderName: 'traits-mania',
      modName: 'Traits Mania',
      modVersion: '1.0.0',
      supportedGameVersion: '1.7.2',
      modRemoteFileId: '',
      modBuildPath: path.join(config.ck3GameModPath, 'traits-mania').replaceAll(/\\/g, '/'),
      modTags: `{
	"Balance"
	"Decisions"
	"Gameplay"
	"Utilities"
}`,
    },

  },
};
