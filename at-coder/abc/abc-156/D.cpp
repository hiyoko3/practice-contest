#include <bits/stdc++.h>
#define repeat(i, n) for(ll (i) = 0; i < (n); i++)
using namespace std;
typedef long long ll;

int main() {
	ll n, a, b;
	cin >> n >> a >> b;
	vector<ll> table(n, 0), reverse(n, 0);
	ll result = 0;
	ll M = ll(1e9) + 7;

	table[0] = 1;
	for(ll i = 1; i < n; i++) {
		table[i] = (table[i-1] * i) % M;
	}

	reverse[n-1] = 1;
	for(ll i = n-2; i >= 0; i--) {
		reverse[i] = (i / reverse[i+1]) % M; // MOD - inv[MOD%i] * (MOD / i) % MOD;
		cout << reverse[i] << endl;
	}

	for(ll i = 0; i < n; i++) {
		result += (table[i] / reverse[i]) % M;
		cout << result << endl;
	}

	result -= (table[a]/ reverse[a]) % M;
	result -= (table[b]/ reverse[b]) % M;

	cout << result << endl;
	return 0;
}
