const path = require('path');

// Assuming 'config' is another module you have that provides the required configurations
const config = require('./')();

function modConfig() {
  const combinedModName = 'mojo';
  const modBuildPath = path.join(config.GameCustomModPath, combinedModName);

  // Define the structure of your combined mod and submods
  return {
    CombinedMod: {
      Replacements: {
        modFolderName: combinedModName,
        modName: 'Mojo by Antick',
        modVersion: '0.1.0',
        supportedGameVersion: '1.11.0.1',
        modRemoteFileId: '',
        modBuildPath: modBuildPath.replace(/\\/g, '/'),
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
    },
    SubMods: {
      'age-of-invasion': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_age_of_invasion',
          modFolderName: 'age-of-invasion',
          modName: 'Age of Invasion',
          modVersion: '1.0.0',
          supportedGameVersion: '1.10.0.1',
          modRemoteFileId: '',
          modBuildPath: path.join(config.GameCustomModPath, 'age-of-invasion').replace(/\\/g, '/'),
          modTags: `{
	"Gameplay"
	"Utilities"
	"Warfare"
}`,
        },
      },
      'grand-council': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_grand_council',
          modFolderName: 'grand-council',
          modName: 'Grand Council',
          modVersion: '1.0.0',
          supportedGameVersion: '1.10.0.1',
          modRemoteFileId: '',
          modBuildPath: path.join(config.GameCustomModPath, 'grand-council').replace(/\\/g, '/'),
          modTags: `{
	"Gameplay"
	"Utilities"
}`,
        },
      },
      'refurbished-titles': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_refurbished_titles',
          modFolderName: 'refurbished-titles',
          modName: 'Refurbished Titles',
          modVersion: '1.0.0',
          supportedGameVersion: '1.10.0.1',
          modRemoteFileId: '',
          modBuildPath: path.join(config.GameCustomModPath, 'refurbished-titles').replace(/\\/g, '/'),
          modTags: `{
	"Culture"
	"Fixes"
}`,
        },
      },
      'tweak-n-treat': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_tweak_it',
          modFolderName: 'tweak-n-treat',
          modName: 'Tweak n Treat',
          modVersion: '1.0.0',
          supportedGameVersion: '1.10.0.1',
          modRemoteFileId: '',
          modBuildPath: path.join(config.GameCustomModPath, 'tweak-n-treat').replace(/\\/g, '/'),
          modTags: `{
	"Fixes"
}`,
        },
      },
      'united-thrones': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_united_thrones',
          modFolderName: 'united-thrones',
          modName: 'United Thrones',
          modVersion: '1.0.0',
          supportedGameVersion: '1.10.0.1',
          modRemoteFileId: '',
          modBuildPath: path.join(config.GameCustomModPath, 'united-thrones').replace(/\\/g, '/'),
          modTags: `{
	"Fixes"
}`,
        },
      },
      'let-there-be-music': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_let_there_be_music',
          modFolderName: 'let-there-be-music',
          modName: 'Let there be music',
          modVersion: '1.0.0',
          supportedGameVersion: '1.10.0.1',
          modRemoteFileId: '',
          modBuildPath: path.join(config.GameCustomModPath, 'let-there-be-music').replace(/\\/g, '/'),
          modTags: `{
	"Utilities"
}`,
        },
      },
      'auto-pause-game': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_auto_pause',
          modFolderName: 'auto-pause-game',
          modName: 'Auto Pause Game',
          modVersion: '1.0.1',
          supportedGameVersion: '1.11.0.1',
          modRemoteFileId: 'remote_file_id="2906586207"',
          modBuildPath: path.join(config.GameCustomModPath, 'auto-pause-game').replace(/\\/g, '/'),
          modTags: `{
	"Fixes"
	"Utilities"
}`,
        },
      },
      'epe-barbershop-idle-pose-compatch': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_idle_pose_compatch',
          modFolderName: 'epe-barbershop-idle-pose-compatch',
          modName: 'EPE + Barbershop Idle Pose Compatch By Antick',
          modVersion: '1.1.1',
          supportedGameVersion: '1.8.1',
          modRemoteFileId: 'remote_file_id="2870777363"',
          modBuildPath: path.join(config.GameCustomModPath, 'epe-barbershop-idle-pose-compatch').replace(/\\/g, '/'),
          modTags: `{
	"Fixes"
	"Portraits"
}`,
        },
      },
      'traits-mania': {
        Enabled: true,
        Replacements: {
          modId: config.ModIdPrefix + '_traits_mania',
          modFolderName: 'traits-mania',
          modName: 'Traits Mania',
          modVersion: '1.0.0',
          supportedGameVersion: '1.7.2',
          modRemoteFileId: 'remote_file_id="2870777363"',
          modBuildPath: path.join(config.GameCustomModPath, 'traits-mania').replace(/\\/g, '/'),
          modTags: `{
	"Balance"
	"Decisions"
	"Gameplay"
	"Utilities"
}`,
        },
      },
      // ... more submods ...
    },
  };
}

module.exports = modConfig;
