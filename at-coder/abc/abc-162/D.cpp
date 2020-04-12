#include <bits/stdc++.h>
#include <boost/multiprecision/cpp_int.hpp>

// increment
#define rep_n(_1, _2, _3, NAME, ...) NAME
#define rep_2(i, n) for (int (i) = 0; (i) < (int)(n); ++(i))
#define rep_3(i, initial, n) for (int (i) = (int)(initial); (i) < (int)(n); ++(i))
#define repeat(...) rep_n(__VA_ARGS__, rep_3, rep_2, rep_1)(__VA_ARGS__)
// decrement
#define rep_r_n(_1, _2, _3, NAME, ...) NAME
#define rep_r_2(i, n) for (int (i) = (int)(n); (i) > 0; (i)--)
#define rep_r_3(i, initial, n) for (int (i) = (int)(initial); (i) > (int)(n); (i)--)
#define repeat_r(...) rep_r_n(__VA_ARGS__, rep_r_3, rep_r_2)(__VA_ARGS__)
// define variable
using namespace std;
typedef long long ll;
using mpint = boost::multiprecision::cpp_int; // 多倍長整数（任意精度）
typedef pair<int, int> P;
const double EPS = numeric_limits<double>::epsilon();
// 最大公約数
ll gcd(ll a, ll b) { return b ? gcd(b, a % b) : a; }
// 最小公倍数
ll lcm(ll a, ll b) { return a / gcd(a, b) * b; }
// char -> int
int ctoi(char c) { return (int)(c - '0'); }
// 素因数分解
vector< pair<ll, ll> > prime_factorize(ll n) {
	vector< pair<ll, ll> > p;
	for(ll i = 2; i*i <= n; i++) {
		if ((n % i) != 0) continue;
		p.push_back(make_pair(i, 0));
		while((n % i) == 0) {
			n /= i;
			p[i].second += 1;
		}
		p.push_back(make_pair(n, 0));
	}

	if (n != 1) p.push_back(make_pair(n, 1));
	return p;
}
string alphabet() { return "abcdefghijklmnopqrstuvwxyz"; }

int main() {
	int N;
	string S;
	cin >> N >> S;

	vector<int> r, g, b;
	repeat(i, S.length()) {
		string s = S.substr(i, 1);

		if (s == "R") r.push_back(i+1);
		else if (s == "G") g.push_back(i+1);
		else if (s == "B") b.push_back(i+1);
	}

	ll result = r.size() * g.size() * b.size();
	for(int i = 0; i < N; i++){
		for(int j = i+1; j < N; j++){
				int k = (2*j) - i;

				if(k >= N) break;

				if((S[i] != S[j]) && (S[j] != S[k]) && (S[i] != S[k])) result--;
		}
	}

	cout << result << endl;
	return 0;
}