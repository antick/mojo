# CK3 Mojo

CK3 Mojo is a versatile tool designed for creating and managing mods for CK3 (Crusader Kings 3).

## What is CK3 Mojo?

CK3 Mojo enables you to develop mods as separate modules. This approach facilitates easy management and modular 
development. When building your mods, you can choose which modules to include, allowing you to create larger and 
more intricate mods.

By breaking down complex mods into submods, you can maintain compatibility, manage updates for new CK3 versions, and 
avoid concerns about load order. Refer to the documentation to learn more about writing submods, building comprehensive 
mods, and even publishing submods individually on Steam – all according to your preferences.

## Installation

Install Go 1.20 or higher and Node.js 18 or higher to setup the development environment.

## Usage

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Mojo Mod

The Mojo tool not only serves as a mod builder but also empowers you to simultaneously develop multiple submods. This 
functionality allows you to construct your own comprehensive overhaul mod. Instead of cluttering your mod list with 
numerous entries and worrying about compatibility, you can consolidate them as submods here. Even if one mod encounters 
issues, you can still proceed with building other mods within the final mod folder. If utilized effectively, this tool 
is very powerful for creating CK3 mods.

### Load order

Place this mod at the bottom of your load order to avoid overriding other mods. However, you can adjust this based on 
your preferences, especially if you intend to avoid interfering with other mods.

### Compatibility

This mod is Ironman compatible and supports Achievements. Note that individual submods may exhibit varying 
compatibility, so review the submod list for comprehensive insights.

Compatibility with other mods may vary. A list of supported and unsupported mods will be updated as users provide 
feedback through the mod's Steam page or GitHub issues.

Given the substantial changes in this mod, compatibility with existing saved games may be affected. Starting a new 
game might be necessary for optimal functionality. The impact depends on the submods included, so please report any 
encountered issues to help improve the mod.

### Supported mods

To be updated.

### Unsupported mods

To be updated.

## Features

1. Conversion of paradox script files to JSON format
2. Collection of pre-generated male and female DNA
3. Collection of COA (Coat-of-Arms) templates

## Planned features

1. New and intriguing events involving friends, family, enemies, vassals, courtiers, councilors, and your realm, enhancing gameplay diversity.
2. Family reunion and diverse religious events on cultural, religious, and regional occasions.
3. Expansion of traits and associated events.
4. Additional decisions to enrich gameplay options.
5. Events for skill advancement.
6. Rare and distinctive events offering opportunities to acquire resources like gold, prestige, piety, and renown.
7. Raiding events.
8. Trading events and decisions.

## Known issues

## FAQ

## Support

This project is a labor of love during my free time. I'll do my best to provide support whenever possible.
