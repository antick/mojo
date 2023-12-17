const readline = require('readline');
const app = require('./src/app');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

function main() {
  console.log('----------------');
  console.log('Welcome to Mojo!');
  console.log('----------------');
  console.log('Please enter your choice:');
  console.log('1: Build Mod (Local)');
  console.log('2: Build Mod (Game)');
  console.log('3: Pull CK3 Game Files');
  console.log('\nYour choice (1/2/3): ');

  rl.question('', answer => {
    answer = answer.trim();
    switch (answer) {
      case '1':
        app.buildModsInLocal([], true)
          .then(() => console.log('Mods built successfully in local!'))
          .catch((err) => console.error('Error building local mods:', err));
        break;
      case '2':
        app.buildModsInGame([], true)
          .then(() => console.log('Mods built successfully in game folder!'))
          .catch((err) => console.error('Error building game mods:', err));
        break;
      case '3':
        app.syncCk3GameFiles()
          .then(() => console.log('CK3 game files synced successfully!'))
          .catch((err) => console.error('Error syncing CK3 game files:', err));
        break;
      default:
        console.error('Invalid choice');
        break;
    }

    rl.close();
  });
}
