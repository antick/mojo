const readline = require('readline');
const {exec} = require('child_process');
const config = require('./src/config')

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
    answer = answer.trim()
    switch (answer) {
      case '1':
        buildModsInLocal([], true)
          .then(() => console.log('Mods built successfully!'))
          .catch((err) => console.error('Error building mods:', err));
        break;
      case '2':
        buildModsInGame([], true)
          .then(() => console.log('Mods built successfully!'))
          .catch((err) => console.error('Error building mods:', err));
        break;
      case '3':
        pullCk3GameFiles()
          .then(() => console.log('CK3 game files pulled successfully!'))
          .catch((err) => console.error('Error pulling CK3 game files:', err));
        break;
      default:
        console.error("Invalid choice");
        break;
    }

    rl.close();
  });

}

