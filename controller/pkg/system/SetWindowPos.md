# SetWindowPos

## hWnd	

- HWND	类型。
- 要定位的窗口的句柄。

## hWndInsertAfter

- HWND 类型。
- 在 Z 顺序中位于定位窗口之前的窗口的句柄。

此参数可以是窗口句柄，也可以是以下值之一：

- HWND_BOTTOM (1): 将窗口置于 Z 顺序的底部。 
- HWND_NOTOPMOST (-2): 将窗口置于所有最顶层窗口的下方（即在所有非最顶层窗口的顶部）。
- HWND_TOP (0): 将窗口置于 Z 顺序的顶部。
- HWND_TOPMOST (-1): 将窗口置于所有非最顶层窗口之上，并保持在此位置。

## X	

- int	
- 窗口左上角的新 X 坐标。

## Y	

- int	
- 窗口左上角的新 Y 坐标。

## cx	

- int	
- 窗口的新宽度。
## cy	

- int	
- 窗口的新高度。


## uFlags

- UINT 类型。
- 窗口大小和位置标志。

此参数可以是以下一个或多个值的组合：

- SWP_NOSIZE (0x0001): 保持当前大小（忽略 cx 和 cy 参数）。
- SWP_NOMOVE (0x0002): 保持当前位置（忽略 X 和 Y 参数）。
- SWP_NOZORDER (0x0004): 保持当前的 Z 顺序（忽略 hWndInsertAfter 参数）。
- SWP_NOREDRAW (0x0008): 不重绘更改。当设置此标志时，不会发生任何类型的重绘。
- SWP_NOACTIVATE (0x0010): 不激活窗口。
- SWP_SHOWWINDOW (0x0040): 显示窗口。
- SWP_HIDEWINDOW (0x0080): 隐藏窗口。