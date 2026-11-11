#!/usr/bin/env bash

sleep 0.5

# Sets a random wallpaper with hyprpaper
wallpapers=($(ls -d $HOME/Pictures/wallpapers/*.png))
wallpapers+=($(ls -d $HOME/Pictures/wallpapers/*.jpg))
#wallpapers+=($(ls -d $HOME/Pictures/wallpapers/*.jpeg))
#wallpapers+=($(ls -d /usr/share/hyprland/wall*))

wall=${wallpapers[ $RANDOM % ${#wallpapers[@]} ]}

hyprctl hyprpaper wallpaper ,$wall
