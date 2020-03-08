<?php
while ( ! feof(STDIN) ) {
  $line = trim(fgets(STDIN));
  echo "${line}\n";
}
