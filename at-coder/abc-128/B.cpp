#include <bits/stdc++.h>
#define repeat(i, n) for(ll i = 0; i < (n); i++)
using namespace std;
typedef long long ll;
typedef pair<string, int> book;

bool compare(pair<string, int> a, pair<string, int> b) {
	// 基本はfirstで比較
    if(a.first != b.first){
        return a.first < b.first; // 昇順
    }

    return a.second > b.second; // 降順
}


int main() {
	int N = 0;
	cin >> N;
	vector<book> books(N);
	vector<int> indexes(N);

	repeat(i, N) {
		string city = "";
		int points = 0;
		cin >> city >> points;
		books[i].first = city;
		books[i].second = points;
		indexes[i] = points;
	}

  // sort(books.begin(), books.end());
  sort(books.begin(), books.end(), compare);
	repeat(i, N) {
		int point = books[i].second;
		auto itr = find(indexes.begin(), indexes.end(), point);
		int index = distance(indexes.begin(), itr);
		cout << index + 1 << endl;
	}
	return 0;
}
