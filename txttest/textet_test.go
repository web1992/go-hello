package txttest

import "testing"

// for i in `cat r.txt |awk -F "\r" '{print $1}' `;do  echo $i; curl $i>>resp.txt; echo "\n" >> resp.txt ;done;
func TestCurl(t *testing.T) {
	Curl()
}
