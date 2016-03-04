APIs:
Get /user/login?token=&name=&userid

redis exec
MULTI
SADD users userid
HMSET userid token token name name
EXEC

Get /user/id/
HGETALL userid

GET /wallpapers/fromuser?user_id

获取用户的所有照片
GET /wallpapers/list?

GET /wallpaper/download?wallpaperid
下载对应id
HGET wallpaper:id url

Post /wallpaper/upload?userid


return json 

GET /likes? wallpaperid 
获取某个壁纸的点赞数
POST /like?wallpaperid=&userid=
userid为某个壁纸点赞

