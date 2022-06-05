import os
import shutil
import re

path = input("Enter the path : ")
files = [f for f in os.listdir(path) if os.path.isfile(os.path.join(path, f))]

def get_season(filename:str):
    pattern_regex = r"([Ss](\d+))"
    a=re.findall(pattern_regex, filename)
    if len(a) > 0:
        return "S"+str("{:02d}".format(int(a[0][1])))
    else:
        return None

def get_resolution(filename:str):
    pattern_regex = r"((480|1080|720|2160)[Pp])"
    a=re.findall(pattern_regex, filename)
    if len(a) > 0:
        return a[0][0].upper()
    else:
        return None

def get_encode(filename:str):
    text = ""
    pattern_regex = r"([Xx]265)"
    if re.search(pattern_regex,filename) != None:
        text+="X265"
    if re.search(r"(10.?[Bb][Ii][Tt])",filename) or re.search(r"[Hh][Ee][Vv][Cc]",filename) != None:
        text+=" 10bit" if text != "" else "10bit"
    else:
        return None
    return text

def get_type(filename:str):
    if re.search(r"[Dd][Uu][Bb][Bb]?[Ll]?[Ee][Dd]?",filename) != None:
        return "Dubbled"
    elif re.search(r"[Ss][Oo|Uu][Ff|Bb][Tt]?[Bb]?[Ee]?[Dd]?",filename) != None:
        return "SoftSub"
    else:
        return None

def get_file_path(filename:str):
    newpath = path 
    season = get_season(filename)
    if season != None:
        newpath += f"/{season}"
    type = get_type(filename)
    if type != None:
        newpath += f"/{type}"
    resolution = get_resolution(filename)
    if resolution != None:
        newpath += f"/{resolution}"
    encodes= get_encode(filename)
    if encodes != None:
        newpath += f" {encodes}"

    return newpath
    
def create_folder_season(files:list):
    fol =[]
    for i in files:
        season = get_season(i)
        if season != None:
            if season not in fol:
                fol.append(season)
                new = path+f"/{season}"
                if os.path.exists(new) != True:
                    os.mkdir(new)

def create_folder_type(files:list):
    for i in files:
        type = get_type(i)
        if type != None:
            season = get_season(i)
            if season != None:
                    new = path+f"/{season}/{type}"
                    if os.path.exists(new) != True:
                        os.mkdir(new)

def create_folder_resolution(files:list):
    newpath = ""
    for i in files:
        resolution = get_resolution(i)
        encodes= get_encode(i)
        if resolution != None:
            newpath = path+f"/{resolution}"
        if encodes != None:
            resolution += f" {encodes}"

        encodes= get_encode(i)
        if encodes != None:
            newpath += f" {encodes}"

        season = get_season(i)
        if season != None:
            gtype= get_type(i)
            if gtype != None:
                season += f"/{gtype}"   
            new = path+f"/{season}/{resolution}"
            if os.path.exists(new) != True:
                os.mkdir(new)

def move_files(files:list):
    for i in files:
        newpath = get_file_path(i)
        if os.path.exists(newpath) == True:
            shutil.move(path+f"/{i}",newpath+f"/{i}")

create_folder_season(files)
create_folder_type(files)
create_folder_resolution(files)
move_files(files)
