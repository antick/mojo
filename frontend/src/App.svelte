<script lang="ts">
  import './style.css';
  import { Icon } from 'svelte-awesome-icons';
  import {
    BuildModsInLocal,
    BuildModsInGame,
    PullCk3GameFiles,
    GetModList,
  } from '../wailsjs/go/main/App.js'

  let activeTab = 'tab1';
  let resultText: string = ""
  let buildingInProgress: boolean = false
  let checkedMods = new Set();
  let buildCombinedMod: boolean = false

  $: subMods = null

  $: buttonClasses = buildingInProgress
    ? "border rounded btn bg-gray-300 text-gray-700 p-2 cursor-not-allowed"
    : "border rounded btn bg-gray-300 text-gray-700 p-2"

  function buildModsInLocal(): void {
    const selectedModKeys = Array.from(checkedMods)

    buildingInProgress = true
    resultText = "Building mods in local..."
    BuildModsInLocal(selectedModKeys, buildCombinedMod)
      .then(() => {
        resultText = "Build generated in local build folder"
        buildingInProgress = false
      })
      .catch((err) => {
          resultText = err
          buildingInProgress = false
      })
  }

  function buildModsInGame(): void {
    const selectedModKeys = Array.from(checkedMods)

    buildingInProgress = true
    resultText = "Building mods in game..."
    BuildModsInGame(selectedModKeys, buildCombinedMod)
      .then(() => {
        resultText = "Build generated in game build folder"
        buildingInProgress = false
      })
      .catch((err) => {
        resultText = err
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

  function handleModSelection(modKey, event) {
    if (event.target.checked) {
      checkedMods.add(modKey);
    } else {
      checkedMods.delete(modKey);
    }
  }

  function handleModCombinationSettings(event) {
    buildCombinedMod = !!event.target.checked;
  }

  // Fetch the mod list from the backend on page load
  GetModList().then((modList: any) => {
    subMods = modList.SubMods
  })
</script>

<main>
  <div class="flex">
    <div class="w-1/4 border border-gray-300 h-screen">
      <div class="p-6 flex justify-center font-bold">
        Mojo Mod Builder
      </div>
      <div class="">
        <button class="{activeTab === 'tab1' ? 'bg-teal-400' : 'bg-gray-200'} border-y border-gray-300 w-full p-6 text-left hover:bg-teal-400 focus:outline-none focus:bg-teal-400" on:click={() => (activeTab = 'tab1')}>
          <span class="flex items-center gap-1 ml-1"><Icon name="hammer-solid" class="w-4" /> Builder</span>
        </button>
        <button class="{activeTab === 'tab2' ? 'bg-teal-400' : 'bg-gray-200'} border-b border-gray-300 w-full p-6 text-left hover:bg-teal-400 focus:outline-none focus:bg-teal-400" on:click={() => (activeTab = 'tab2')}>
          <span class="flex items-center gap-1 ml-1"><Icon name="angles-down-solid" class="w-4" /> Puller</span>
        </button>
        <button class="{activeTab === 'tab3' ? 'bg-teal-400' : 'bg-gray-200'} border-b border-gray-300 w-full p-6 text-left hover:bg-teal-400 focus:outline-none focus:bg-teal-400" on:click={() => (activeTab = 'tab3')}>
          <span class="flex items-center gap-1 ml-1"><Icon name="file-pen-solid" class="w-4" /> Fixer</span>
        </button>
      </div>
    </div>

    <div class="w-3/4 flex flex-col gap-2 p-10">
      <div>
        {#if resultText !== ""}
          <div class="max-w p-1 mb-4 bg-gray-100 border border-gray-200 rounded-lg shadow">
            <div class="px-4 text-teal-600 text-sm">{resultText}</div>
          </div>
        {/if}
        {#if activeTab === 'tab1'}
          <div class="max-w p-6 bg-gray-100 border border-gray-200 rounded-lg shadow">
            <p class="flex items-center mb-4">
              Select the mods that you want to build:
            </p>

            {#if subMods}
              {#each Object.entries(subMods) as [modKey, subMod]}
                <div class="flex items-center mt-2">
                  <input
                    type="checkbox"
                    value=""
                    class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
                    on:change={event => handleModSelection(modKey, event)}
                  >
                  <label for="default-checkbox" class="ms-2 text-sm font-medium text-gray-900" id="{modKey}"> {subMod.Replacements.modName}</label>
                </div>
              {/each}
            {/if}

            <p class="flex items-center my-4">
              Do you want to build a combined mod?
            </p>

            <div class="flex items-center mt-2">
              <input
                type="checkbox"
                value=""
                class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
                on:change={event => handleModCombinationSettings(event)}
              >
              <label for="default-checkbox" class="ms-2 text-sm font-medium text-gray-900"> Yes, create a single (combined) mod from the selected choices mentioned above.</label>
            </div>

            <div class="flex gap-6 mt-6">
              <button class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-teal-500 rounded-lg hover:bg-teal-600 focus:outline-none" on:click={buildModsInLocal} disabled={buildingInProgress}>
                Build Mods (Local)
              </button>

              <button class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-teal-500 rounded-lg hover:bg-teal-600 focus:outline-none" on:click={buildModsInGame} disabled={buildingInProgress}>
                Build Mods (Game)
              </button>
            </div>
          </div>

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
  </div>
</main>
