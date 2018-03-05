-- #df:changeUser(system)#
-- #df:checkUser(mainSchema)#
create database /*$dfprop.mainCatalog*/ character set utf8;

-- #df:reviveUser()#
-- #df:checkUser(mainUser, grant)#
grant all privileges on /*$dfprop.mainCatalog*/.*
  to /*$dfprop.mainUser*/@localhost identified by '/*$dfprop.mainPassword*/';

flush privileges;
