APIs:
Get /user/login?token=&name=&userid

redis exec
MULTI
userid := email@facebook
SADD users userid
HMSET user:id token token name name
EXEC

Get /user/id/
HGETALL userid

GET /wallpapers/fromuser?user_id
SMEMBERES user:id:wp
获取用户的所有照片

GET /wallpapers/list?
sorted set: wallpapers

GET /wallpaper/download?wallpaperid
下载对应id
HGET wallpaper:id url

Post /wallpaper/upload?userid
wallpaperid := md5 wallpaper
ZADD wallpapers wallpaperid
HMSET wallpaper:id userid
SADD user:id:w wallpaperid

return json 

GET /likes? wallpaperid 
获取某个壁纸的点赞数
POST /like?wallpaperid=&userid=
userid为某个壁纸点赞

