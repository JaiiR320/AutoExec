# AutoExec: A simple CLI for auto booting apps

## Description

AutoExec was inspired by CS:GO, where players would write a file containing in game commands that are run upon startup. These commands were used to customize the settings for the player automatically, which is especially useful if you have dozens of settings that are fine tuned. The advantage of this was that it was a single file that was portable, meaning it can be added to any instance of CS:GO, automatically implementing the settings that you so tedeously tweaked.

This version is slightly different, but has the same vision; automatically launching apps and tweaking settings on launch. 

> Still a WIP so some paths are hard coded. It might not work on your PC!

## Installation

To install AutoExec, simply add the executable to a directory on your PC, and add the executable to your PATH environment variable for easy access. 

## Usage

Using the AutoExec app requires installing the CLI. Afterwhich, the user can interact with the CLI to add a file called a "bucket" which holds all commands and applications to launch. You can create any number of buckets for different situations such as a gaming bucket which launches your games and Discord, or a production bucket, lanuching VSCode and Firefox. It is up to you on which apps and settings are ran. 

Typing in the console `autoexec help` will print out a list of commands as well as their description. 

## Documentation 

Full documentation on how to use AutoExec can be found [Here](https://auto-exec.vercel.app/)

> also still a work in prograss, might be a default docusaurus project upon visiting...

## Conclusion

AutoExec was built in only a few hours on a random Thursday, so please don't sue me if I brick your PC. Jokes aside, I built this CLI to learn Go and also to have a nice excuse to write some documentation. Feel free to download and explore the code!