# Screen Layout Tool

Screen Layout Tool is a software to help you organize and arrange windows on your computer screen, especially when using multiple monitors or working with multiple applications.

Screen Layout Tool committed to making multi-screen window management more efficient, precise and smooth.

<video controls="" muted="" autoplay="" loop="" width="100%">
    <source src="https://jialo.com/videos/demo.mp4" type="video/mp4">
</video>

## 💎 It is easy to use

**Fast Select Layout Solution**

Use the right-click menu of the tray icon or hotkeys to select a layout solution.

| hotkey                         | description            |
| :----------------------------- | :--------------------- |
| `Win` + `Shift` + `MouseWheel` | Change layout solution |
| `Win` + `Shift` + `←` / `→`    | Change layout solution |

- The Screen Layout Tool has two commonly used layout solutions built-in.
- You can also design your own layouts or import layout solutions shared by others.

**Fast Switch Window Position**

Use the hotkeys to switch the position of the current window in the layout solution.

| hotkey               | description            |
| :------------------- | :--------------------- |
| `Win` + `MouseWheel` | Change window position |
| `Win` + `←` / `→`    | Change window position |

- The first stop for a window is the closest position in layout when Previous/Next be used.
- Roughly adjust the window's size and position to your desired, then switch, It's very likely what you want.

## 💎 You can import layouts you like

You can import layout solutions shared by others.

- Download your favorite layout file from [official website](https://jialo.com/layouts.php) or [official community](https://www.reddit.com/r/ScreenLayoutTool/).
- Click "Open Layout Folder" in the right-click menu of the tray icon to put the layout file in it.
- Click "Reload", the layout solution will appear in "Layout" menu.

## 💎 You can easily design your layout

You can also design your own screen layout solution using the `.json` file.

```json
{
  "author": "Tyx",
  "name": "Default",
  "version": "1.0",
  "description": "Default layout solution",
  "positions": [
    {
      "left": 0,
      "width": 20,
      "top": 0,
      "height": 100,
      "monitor": 0
    },
    {
      "right": 0,
      "width": 20,
      "top": 0,
      "height": 100,
      "monitor": 1
    },
    {
      "top": 20,
      "bottom": 20,
      "left": 20,
      "right": 20,
      "monitor": 2
    }
  ]
}
```

| no. | field   | unit | range | description                 |
| :-: | ------- | :--: | :---: | --------------------------- |
|  1  | height  |  %   | 0-100 | Height of the window        |
|  2  | top     |  %   | 0-100 | Distance from edge to edge  |
|  3  | bottom  |  %   | 0-100 | Distance from edge to edge  |
|  4  | width   |  %   | 0-100 | Width of the window         |
|  5  | left    |  %   | 0-100 | Distance from edge to edge  |
|  6  | right   |  %   | 0-100 | Distance from edge to edge  |
|  7  | monitor |      | 0-100 | Monitor number, 0: primary. |

- 1-3 must choose 2.
- 3-4 must choose 2.
- 7 is optional, the default is 0.

***For more information please visit the official website [jialo.com](https://jialo.com).***