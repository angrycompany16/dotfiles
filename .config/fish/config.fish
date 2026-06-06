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







	#	function tash --description "Simple command that moves the file into the current directory"
	# set file "$argv[1]"
	# sudo mv "$file" "./"(basename "$file")
	#end

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

	function ga --description "git add shortcut"
		command git add $argv
	end

	function gc --description "git commit shortcut"
		command git commit $argv
	end

	function gp --description "git push shortcut"
		command git push $argv
	end
end
