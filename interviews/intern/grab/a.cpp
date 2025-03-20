#include <bits/stdc++.h>
#include <string>
#include <iostream>

using namespace std;

int solution(string &s) {
    int n = s.length();
    if (n < 3) return 0;

    // The maximum number of patterns we can form is n/3
    int max_patterns = n / 3;
    
    // Count the frequency of each character
    vector<int> freq(26, 0);
    for (char c : s) freq[c - 'a']++;
    
    // Greedy approach:
    // 1. First try to form "aba" patterns (where a and b are different)
    // 2. Then use remaining characters to form "aaa" patterns
    
    int total = 0;
    
    // Try to form "aba" patterns first
    for (int i = 0; i < 26; i++) {
        for (int j = 0; j < 26; j++) {
            if (i == j) continue; // Skip if characters are the same
            
            // We need 2 of character i and 1 of character j
            int patterns = min(freq[i] / 2, freq[j]);
            total += patterns;
            
            freq[i] -= patterns * 2;
            freq[j] -= patterns;
        }
    }
    
    // Use remaining characters to form "aaa" patterns
    for (int i = 0; i < 26; i++) {
        total += freq[i] / 3;
    }
    
    // Debug print to trace values
    cout << "Total patterns found: " << total << endl;
    
    return min(total, max_patterns);
}

int main() {
    string s1 = "aaabac";
    cout << "Test 1: " << (solution(s1) == 2 ? "true" : "false") << endl;
    
    string s2 = "xyzwyz";
    cout << "Test 2: " << (solution(s2) == 2 ? "true" : "false") << endl;
    
    string s3 = "dd";
    cout << "Test 3: " << (solution(s3) == 0 ? "true" : "false") << endl;
    
    string s4 = "fknfknf";
    cout << "Test 4: " << (solution(s4) == 2 ? "true" : "false") << endl;

    string s5 = "ddddddddd";
    cout << "Test 5: " << (solution(s5) == 3 ? "true" : "false") << endl;
    
    return 0;
}