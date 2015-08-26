# localize
Creates localized versions of all kind of files.

Usage: localize input-folder localization-folder output-folder

directory structure
.
├── input
│   └── index.html
└── localization
    ├── files.es
    ├── like.it
    ├── name.de
    ├── want.pl
    └── you.no

You can name the files inside the localization-folder like you want.
Only the file extension is important.
But don´t use '.' in your filename as the '.' is used to identify the extension!

If you now run 'localize input localization output' you get the following directory structure:
.
├── input
│   └── index.html
├── localization
│   ├── files.es
│   ├── like.it
│   ├── name.de
│   ├── want.pl
│   └── you.no
└── output
    ├── de
    │   └── index.html
    ├── es
    │   └── index.html
    ├── index.html
    ├── it
    │   └── index.html
    ├── no
    │   └── index.html
    └── pl
        └── index.html
