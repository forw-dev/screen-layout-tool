# Screen Layout Tool Controller

## window

### get

get current window info

### move

- 操作对象是窗口整体。
- 尺寸固定。

参数

- `--top`
- `--bottom`
- `--left`
- `--right`
- `--monitor`
  - 0 `primary`
- `--locate`
  - `0` offset，自身定位，-100~100
  - `1` coordinate，坐标定位，0~100
  - `2` margin，边缘定位，0~100

样例

`cmd window move -top nnn --offset`

### pull

- 操作对象是窗口的边。
- 对边固定。

参数

- `--top`
- `--bottom`
- `--left`
- `--right`
- `--locate`
  - `0` offset，自身定位，-100~100
  - `1` coordinate，坐标定位，0~100
  - `2` margin，边缘定位，0~100

样例

`cmd window pull -top nnn --offset`

### resize

- 操作对像是边的尺寸。
- 中点固定。

参数

- `--width`
- `--height`
- `--locate`
  - `offset` 自身定位 -100~100
  - `fix` 固定值 0~100

### scroll

- `--layout=layout.json`
- `--cross`
- `--step` -any~any，none 最近位置

### jump

- `--layout=layout.json`
- `--index 0` 0~n，none 最近位置

## monitor

get monitors info