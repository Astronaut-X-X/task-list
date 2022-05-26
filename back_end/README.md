# TASKLIST back_end

## model

**user**  
> username (6-16 [0-9a-zA-Z_\-]{6,16})   
> password (6-16 [0-9a-zA-Z@#$%^&*]{6,16})   
> image  (/static/userimag usernametoHash.png)
> email  (/^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$/)
> gorm.model

**task**  
> userid (check user is isExit)  
> done (uint 0 1 10 11)  
> cotent (string)  
> time (2006-01-02)  
> meridiem ("am"|"pm")  
> color (#000000)  
> gorm.model  

**todo**
> userid (check user is isExit)  
> done (uint 0 1 10 11)  
> cotent (string)  
> time (2006-01-02)  
> gorm.model  

**gold**
> userid (check user is isExit)  
> done (uint 0 1 10 11)  
> cotent (string)  
> time (2006-01-02)  
> gorm.model  

## TODO

-[ ] email register  
-[ ] jwt  

## module

github.com/jordan-wright/email
