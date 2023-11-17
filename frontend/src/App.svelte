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

  $: subMods = null

  $: buttonClasses = buildingInProgress
    ? "border rounded btn bg-gray-300 text-gray-700 p-2 cursor-not-allowed"
    : "border rounded btn bg-gray-300 text-gray-700 p-2"

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
              List of Available Mods
              <svg class="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
              </svg>
            </p>
            {#if subMods}
              {#each Object.entries(subMods) as [modKey, subMod]}
                <div class="flex items-center mt-2">
                  <input id="default-checkbox" type="checkbox" value="" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2">
                  <label for="default-checkbox" class="ms-2 text-sm font-medium text-gray-900" id="{modKey}"> {subMod.Replacements.modName}</label>
                </div>
              {/each}
            {/if}

            <p class="mt-4 font-normal text-sm text-gray-700">
              Select the mods that you want to build. You can either select all of them or choose any specific mod. If you select only one mod, it will create a
              mod folder for that particular selection. However, if you select more than one mod, it will create a combined mod called Mojo.
            </p>

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
