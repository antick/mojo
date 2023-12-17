const fs = require('fs');
const path = require('path');

async function readFile(filePath, encoding) {
  return new Promise((resolve, reject) => {
    fs.readFile(filePath, encoding, (err, data) => {
      if (err) {
        reject(err);
      } else {
        resolve(data);
      }
    });
  });
}

async function writeFile(filePath, data) {
  try {
    await fs.promises.writeFile(path.normalize(filePath), data);
  } catch (error) {
    console.error(`Error writing to file: ${error}`);
  }
}

async function updateGameDataPath(filePath, newGameDataPath) {
  try {
    const data = await readFile(filePath, 'utf8');
    const jsonData = JSON.parse(data);

    // Update the gameDataPath property
    jsonData.gameDataPath = newGameDataPath;

    // Write the updated JSON file
    await writeFile(filePath, JSON.stringify(jsonData, null, 2));
    console.log('Game data path successfully updated in launcher-settings.json!');
  } catch (error) {
    console.error(`Error updating game data path: ${error}`);
  }
}

module.exports = {
  updateGameDataPath
};
