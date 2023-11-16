# Mojo

Mojo is a versatile tool designed for creating and managing mods for CK3 (Crusader Kings 3). Please note that this is
under heavy development so a lot of updates would be coming until it reaches to version 1.0.

## What is Mojo?

Mojo lets you make mods in different parts, making it easy to manage and build. With this tool, not only can you pick 
which parts to use for bigger and detailed mods, but you can also work on several small parts at the same time. This 
lets you put together a big mod with all your ideas. Instead of having a lot of separate mods and worrying if they'll 
work well together, you can group them using Mojo. Even if one part has problems, you can keep making the others in the 
main mod space.

By splitting big mods into smaller parts, you can keep them working well, update them for new CK3 versions, and not 
worry about the order they load. Check the guide (WIP) to learn how to write these smaller parts, build full mods, and 
even put them on Steam the way you like.

## Installation

Install Go 1.20 or higher:

https://go.dev/doc/install

Download and install Node.js 20:

https://nodejs.org/en/download

Install Wails from their documentation:

https://wails.io/docs/gettingstarted/installation

## Usage

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## In-Built Mod

This tool also comes up with an in-built mod called Mojo that has a few features to help you get started.

### Load order

Place the in-built mod generated from this tool at the bottom of your load order to avoid overriding other mods. 
However, you can adjust this based on your preferences, especially if you intend to avoid interfering with other mods.

### Compatibility

Compatibility with other mods may vary. A list of supported and unsupported mods will be updated as users provide 
feedback through the mod's Steam page or GitHub issues.

Given the substantial changes in this mod, compatibility with existing saved games may be affected. Starting a new 
game might be necessary for optimal functionality. The impact depends on the submods included, so please report any 
encountered issues to help improve the mod.

### Supported mods

WIP

### Unsupported mods

WIP

## Features

1. Convert pds (paradox script) files to JSON format.
2. Collection of pre-generated male and female DNA.
3. Collection of COA (Coat-of-Arms) templates.
4. In-built collection of mods.

## Planned features

1. New and intriguing events involving friends, family, enemies, vassals, courtiers, councilors, and your realm, enhancing gameplay diversity.
2. Family reunion and diverse religious events on cultural, religious, and regional occasions.
3. Expansion of traits and associated events.
4. Additional decisions to enrich gameplay options.
5. Events for skill advancement.
6. Rare and distinctive events offering opportunities to acquire resources like gold, prestige, piety, and renown.
7. Raiding events.
8. Trading events and decisions.
9. After pulling Ck3 game files, automatically detect any changes that may be required in the mod files you have.
10. Define a temporary path where CK3 game files can be pulled and compared with mod files, eliminating the need for manual comparison.

## Known issues

None at the moment. Please report any issues you encounter.

## FAQ

### Is it Ironman and Achievement compatible?

As of now, yes.

## Support

This free and open-source project is a labor of love created in my free time. I'll do my best to provide support whenever possible.
