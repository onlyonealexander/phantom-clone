# Phantom Clone – School Project

A UI clone of the Phantom crypto wallet app built with plain HTML, CSS, and JavaScript.

## Project Structure

```
phantom-clone/
├── css/
│   └── style.css          # Shared styles
├── js/
│   └── utils.js           # Shared JS utilities
├── pages/
│   ├── index.html         # Home / landing screen
│   ├── import-wallet.html # Import wallet options screen
│   ├── recovery-phrase.html # Enter recovery phrase screen
│   ├── private-key.html   # Enter private key screen
│   └── phantom-site.html  # In-app browser (phantom.com)
└── README.md
```

## How to Run

1. Open the `pages/` folder in VS Code
2. Right-click `index.html` → **Open with Live Server**
   - Or just double-click `index.html` to open in browser
3. Navigate between screens using the buttons

## Screen Flow

```
index.html
  └── Import Wallet → import-wallet.html
        ├── Use Recovery Phrase → recovery-phrase.html
        ├── Use Private Key    → private-key.html
        └── Connect Ledger     → phantom-site.html (in-app browser)
```

## Tech Stack

- Plain HTML5, CSS3, Vanilla JavaScript
- No frameworks or dependencies
- Mobile-first layout (375px phone frame)
