const fs = require('fs');
const path = require('path');

const allowedExtensions = ['.txt', '.info', '.gui', '.yml'];

async function pullCk3Files(mapping) {
  try {
    for (const [srcPath, destPath] of Object.entries(mapping)) {
      await processCk3Directory(srcPath, destPath);
      console.log(`Successfully pulled files from "${srcPath}" to "${destPath}"`);
    }
  } catch (error) {
    console.error(`Error pulling CK3 files: ${error}`);
  }
}

async function processCk3Directory(srcDirectory, destDirectory) {
  // Ensure destination directory exists (create if necessary)
  try {
    await fs.promises.mkdir(destDirectory, { recursive: true });
  } catch (error) {
    if (error.code !== 'EEXIST') {
      throw error;
    }
  }

  const items = await fs.promises.readdir(srcDirectory);

  for (const item of items) {
    const srcPath = path.join(srcDirectory, item);
    const destPath = path.join(destDirectory, item);

    if ((await fs.promises.stat(srcPath)).isDirectory()) {
      // Recursively process subdirectories
      await processCk3Directory(srcPath, destPath);
    } else if (allowedExtensions.includes(path.extname(item))) {
      // Copy files with allowed extensions
      try {
        await fs.promises.copyFile(srcPath, destPath);
      } catch (error) {
        console.error(`Error copying file "${srcPath}": ${error}`);
      }
    }
  }
}
