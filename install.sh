if [ ! -d "$HOME/.config/RSSLauncher" ]; then
    mkdir "$HOME/.config/RSSLauncher"
    echo "//write all your blacklisted words, one word per line" > $HOME/.config/RSSLauncher/blacklist.txt

    echo "//write all your commands with the following format:" > $HOME/.config/RSSLauncher/commands.csv
    echo "//commandName,command (%url will be replaced with the item url)" >> $HOME/.config/RSSLauncher/commands.csv
    echo "onEnter,firefox %url" >> $HOME/.config/RSSLauncher/commands.csv

    echo "" > $HOME/.config/RSSLauncher/feeds.txt
fi
