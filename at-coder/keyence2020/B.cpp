#include <bits/stdc++.h>
using namespace std;
 
int main() {
	int N = 0, result = 0;
	cin >> N;

	int X = 0, L = 0, left = 0, right = 0;
	vector<int> Line;

	for (int i = 0; i < N; i++) {
		cin >> X >> L;
		left = X - L + 1;
		if (left < 0) left = 0;

		right = X + L - 1;
		if (Line[left] == 0 || Line[right] == 0) {
			for (int j = left; j <= right; j++) {
				if (Line[j] == 1) continue;
				Line[j] = 1;
			}
			result++;
		}
	}

	cout << result << endl;
}