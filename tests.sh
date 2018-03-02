# Testing Examples


for file in $(find examples -name \*.sh \
	| grep -v "exit.sh"    \  # thosse three files generate false positive
 	| grep -v "panic.sh"   \  # errors, so its better to exlude them
	| grep -v "signals.sh" \  # from tests.
	); do

	cd $(dirname $file)
	# echo "@" $file
	while IFS='' read -r line || [[ -n "$line" ]]; do
		if [[ "$line" == '$'* ]]; then


			$(printf "$line" | sed -e 's/$ //g' > ./tmp.sh  )
			chmod +x ./tmp.sh
			./tmp.sh > /dev/null 2>&1
			state=$?

			# clean up
			unlink ./tmp.sh

			if (( $state == 1 )); then
				echo "error @ $file @ $line"
				exit 1
			fi
		fi
	done < $(basename $file)

	maybe_binary=$(basename ${file%.*});
	if [[ -a "$maybe_binary" ]]; then
		unlink "$maybe_binary"
	fi

	cd -  > /dev/null 2>&1
done


# Signal and Panic are false positive tests.
