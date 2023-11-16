<script lang="ts">
  import './style.css';
  import { Icon } from 'svelte-awesome-icons';
  import {
    BuildModsInLocal,
    BuildModsInGame,
    PullCk3GameFiles,
  } from '../wailsjs/go/main/App.js'

  let activeTab = 'tab1';
  let name: string
  let resultText: string = ""
  let buildingInProgress: boolean = false

  $: buttonClasses = buildingInProgress
    ? "border rounded btn bg-amber-50 text-gray-700 p-2 cursor-not-allowed"
    : "border rounded btn bg-amber-50 text-gray-700 p-2"

  function buildModsInLocal(): void {
    buildingInProgress = true
    resultText = "Building mods in local..."
    BuildModsInLocal().then(() => {
      resultText = "Build generated in local build folder"
      buildingInProgress = false
    })
  }

  function buildModsInGame(): void {
    buildingInProgress = true
    resultText = "Building mods in game..."
    BuildModsInGame().then(() => {
        resultText = "Build generated in game build folder"
        buildingInProgress = false
    })
  }

  function pullCk3GameFiles(): void {
    buildingInProgress = true
    resultText = "Pulling CK3 files from game folder..."
    PullCk3GameFiles().then(() => {
      resultText = "Successfully pulled CK3 files"
      buildingInProgress = false
    })
  }
</script>

<main>
  <div class="flex">
    <div class="w-1/5 border border-gray-300 h-screen">
      <button class="{activeTab === 'tab1' ? 'bg-teal-400' : 'bg-gray-200'} w-full p-6 text-left hover:bg-teal-400 focus:outline-none focus:bg-teal-400" on:click={() => (activeTab = 'tab1')}>
        <span class="flex items-center gap-1 ml-1"><Icon name="hammer-solid" class="w-4" /> Builder</span>
      </button>
      <button class="{activeTab === 'tab2' ? 'bg-teal-400' : 'bg-gray-200'} w-full p-6 text-left hover:bg-teal-400 focus:outline-none focus:bg-teal-400" on:click={() => (activeTab = 'tab2')}>
        <span class="flex items-center gap-1 ml-1"><Icon name="angles-down-solid" class="w-4" /> Puller</span>
      </button>
      <button class="{activeTab === 'tab3' ? 'bg-teal-400' : 'bg-gray-200'} w-full p-6 text-left hover:bg-teal-400 focus:outline-none focus:bg-teal-400" on:click={() => (activeTab = 'tab3')}>
        <span class="flex items-center gap-1 ml-1"><Icon name="file-pen-solid" class="w-4" /> Fixer</span>
      </button>
    </div>

    <div class="w-3/4 flex flex-col items-center gap-2 mt-10">
      <div class="result" id="result">{resultText}</div>
      {#if activeTab === 'tab1'}
        <div class="p-4">Compile and build your mods in local build folder</div>
        <button class={buttonClasses} on:click={buildModsInLocal} disabled={buildingInProgress}>
          Build Mods (Local)
        </button>
        <div class="p-4">Compile and build your mods in game directly</div>
        <button class={buttonClasses} on:click={buildModsInGame} disabled={buildingInProgress}>
          Build Mods (Game)
        </button>
      {:else if activeTab === 'tab2'}
        <div class="p-4">Copy Game files in provided path for tracking the new changes if CK3 released a new version</div>
        <button class={buttonClasses} on:click={pullCk3GameFiles} disabled={buildingInProgress}>
          Pull CK3 Game Files
        </button>
      {:else if activeTab === 'tab3'}
          Update Launcher Settings
      {/if}
    </div>
  </div>
</main>
