#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
	int numberOfSubstrings(string s){
		int len = s.size();
		int l = 0, r = 0;
		int total = 0;
		vector<int> freq(3,0);

		while(r < len){
			char cur = s[r];
			freq[cur - 'a']++;

			while(validString(freq)){
				total += len - r;
				char leftChar = s[l];
				freq[leftChar - 'a']--;
				l++;
			}
			r++;
		}

		return total;
	}
private:
	bool validString(vector<int>& freq){
		return freq[0] > 0 && freq[1] > 0 && freq[2] > 0;
	}
};

int main(){
	Solution sol;
	string s = "abcabc";
	cout << sol.numberOfSubstrings(s);
	return 0;
}