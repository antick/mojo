<script lang="ts">
  import './style.css';
  import {
    Greet,
    BuildModsInLocal,
    BuildModsInGame,
    PullCk3GameFiles,
    UpdateLauncherSettings,
  } from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter your name below 👇"
  let name: string
  let buildingInProgress: boolean = false

  $: buttonClasses = buildingInProgress
    ? "border rounded btn bg-amber-50 text-gray-700 p-2 cursor-not-allowed"
    : "border rounded btn bg-amber-50 text-gray-700 p-2"

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }

  function buildModsInLocal(): void {
    buildingInProgress = true
    resultText = "Building mods in local..."
    BuildModsInLocal().then(result => {
      resultText = "Build generated in local build folder"
      buildingInProgress = false
    })
  }

  function buildModsInGame(): void {
    buildingInProgress = true
    resultText = "Building mods in game..."
    BuildModsInGame().then(result => {
        resultText = "Build generated in game build folder"
        buildingInProgress = false
    })
  }

  function pullCk3GameFiles(): void {
    buildingInProgress = true
    resultText = "Pulling CK3 files from game folder..."
    PullCk3GameFiles().then(result => {
      resultText = "Successfully pulled CK3 files"
      buildingInProgress = false
    })
  }

  function updateLauncherSettings(): void {
    buildingInProgress = true
    resultText = "Updating launcher settings..."
    UpdateLauncherSettings().then(result => {
      resultText = "Successfully updated launcher settings"
      buildingInProgress = false
    })
  }
</script>

<main class="flex p-10 flex-col justify-center items-center">
  <div class="flex flex-col justify-center items-center gap-2">
    <div class="result" id="result">{resultText}</div>

    <div class="input-box" id="input">
      <input autocomplete="off" bind:value={name} class="input text-gray-700" id="name" type="text"/>
      <button class="border rounded btn bg-amber-50 text-gray-700" on:click={greet}>Greet</button>
    </div>

    <button class={buttonClasses} on:click={buildModsInLocal} disabled={buildingInProgress}>
      Build Mods (Local)
    </button>
    <button class={buttonClasses} on:click={buildModsInGame} disabled={buildingInProgress}>
      Build Mods (Game)
    </button>
    <button class={buttonClasses} on:click={pullCk3GameFiles} disabled={buildingInProgress}>
      Pull CK3 Game Files
    </button>
    <button class={buttonClasses} on:click={updateLauncherSettings} disabled={buildingInProgress}>
      Update Launcher Settings
    </button>
  </div>
</main>

<style>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
