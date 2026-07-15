if status is-interactive
	fastfetch

	function tash --description "Simple command that moves the file into the current directory"
		argparse s/sudo -- $argv
		or return
    
		set file "$argv[1]"
		set dest "./"(basename "$file")

		if set -q _flag_sudo
			sudo mv "$file" "$dest"
		else
			mv "$file" "$dest"	
		end	
	end

	function dup --description "Simple command that copies the file into the current directory"
		argparse s/sudo -- $argv
		or return
    
		set file "$argv[1]"
		set dest "./"(basename "$file")

		if set -q _flag_sudo
			sudo cp -r "$file" "$dest"
		else
			cp -r "$file" "$dest"	
		end	
	end

	function form --description "Create a new folder and cd into it"
		argparse s/sudo -- $argv
		or return
    
		set dir "$argv[1]"

		if set -q _flag_sudo
			sudo mkdir "$dir"
		else
			mkdir "$dir"	
		end	

		cd "$dir"
	end

	function ga --description "git add shortcut"
		command git add $argv
	end

	function gc --description "git commit shortcut"
		command git commit $argv
	end

	function gp --description "git push shortcut"
		command git push $argv
	end

	function docedit --description "Open the typst file in nvim, run typst watch and view it with zathura"
		set file "$argv[1]"
    set pdf (string replace -r '\.typ$' '.pdf' "$file")

		typst watch "$file" >/tmp/typst.log 2>&1 &
		set watch_pid $last_pid

		zathura "$pdf" &
		set zathura_pid $last_pid

		nvim "$file"
		kill $watch_pid $zathura_pid 2>/dev/null
	end
end
