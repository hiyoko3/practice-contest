#include <bits/stdc++.h>
#include <boost/multiprecision/cpp_int.hpp>

// increment
#define rep_n(_1, _2, _3, NAME, ...) NAME
#define rep_2(i, n) for (int(i) = 0; (i) < (int)(n); ++(i))
#define rep_3(i, initial, n) for (int(i) = (int)(initial); (i) < (int)(n); ++(i))
#define rep(...) rep_n(__VA_ARGS__, rep_3, rep_2, rep_1)(__VA_ARGS__)
// decrement
#define rep_r_n(_1, _2, _3, NAME, ...) NAME
#define rep_r_2(i, n) for (int(i) = (int)(n); (i) > 0; (i)--)
#define rep_r_3(i, initial, n) for (int(i) = (int)(initial); (i) > (int)(n); (i)--)
#define rep_r(...) rep_r_n(__VA_ARGS__, rep_r_3, rep_r_2)(__VA_ARGS__)
// define type definition
using namespace std;
typedef long long ll;
using mpint = boost::multiprecision::cpp_int; // 多倍長整数（任意精度）
typedef pair<int, int> P;

// define construct variable
const double EPS = numeric_limits<double>::epsilon();
// 10^9 + 7
const ll MOD = ll(1e9) + 7;
const ll INF = ll(1e10);
const string ALPHABET = "abcdefghijklmnopqrstuvwxyz";
const double PI = acos(-1);

// define methods
// 最大公約数
ll gcd(const ll a, const ll b) { return b ? gcd(b, a % b) : a; }
// 最小公倍数
ll lcm(const ll a, const ll b) { return a / gcd(a, b) * b; }
// char -> int
int ctoi(char c) { return (int)(c - '0'); }
// 素因数分解
vector< pair<ll, ll> > prime_factorize(ll n)
{
	vector<pair<ll, ll>> p;
	for (ll i = 2; i * i <= n; i++)
	{
		if ((n % i) != 0)
			continue;
		pair<ll, ll> a = make_pair(i, 0);
		while ((n % i) == 0)
		{
			n /= i;
			a.second += 1;
		}
		p.push_back(a);
	}

	if (n != 1)
		p.push_back(make_pair(n, 1));
	return p;
}
// 組み合わせの整合性を保つ（後ほど解析する）
ll modinv(ll a)
{
	ll b = MOD, u = 1, v = 0;
	while (b)
	{
		ll t = a / b;
		a -= t * b;
		swap(a, b);
		u -= t * v;
		swap(u, v);
	}
	u %= MOD;
	if (u < 0)
	{
		u += MOD;
	}
	return u;
}
// 順列
ll permutation(const ll n, const ll m)
{
	ll result = 1;
	for (ll i = n; i > n - m; i--)
	{
		result = (result * i) % MOD;
	}
	return result;
}

// 繰り返し n 乗法
ll repeat_square(const ll n, const ll x)
{
	if (x == 0)
		return 1;
	ll res = repeat_square(n, x / 2);
	res = res * res % MOD;
	if ((x % 2) == 1)
		res = res * n % MOD;
	return res;
}
// 組み合わせ nCm
ll combination(ll n, ll m)
{
	ll p = permutation(n, m);
	ll div = 1;
	for (ll i = 2; i <= m; i++)
	{
		div = (div * i) % MOD;
	}
	return (p * modinv(div)) % MOD;
}

void fill(int x, vector<int> &dp) {
	if(dp[x]) return;
	for(int i = 1; i*x <= 1e6; i++) {
		dp[i*x] = true;
	}
	return;
}

int main() {
  int x1, y1;
  int x2, y2;
  int x3, y3;
  cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3;

  double a = sqrt((x1 - x2) * (x1 - x2) + (y1 - y2) *  (y1 - y2));
  double b = sqrt((x2 - x3) * (x2 - x3) + (y2 - y3) *  (y2 - y3));
  double c = sqrt((x1 - x3) * (x1 - x3) + (y1 - y3) *  (y1 - y3));

  double s = (a+b+c) / 2;
  double S = sqrt(s * (s - a) * (s - b) * (s - c));

  cout << fixed << setprecision(1) << S << endl;
	return 0;
}
