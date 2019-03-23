/*  Project 3

    定义一个聚类函数 cluster(int data[], int len, int radius)；
    data中的值没有重复，长度为len,
    把按照数值的聚类进行分为n组，
    对于组G中任意一个数值a,总是能在本组G中找到一个数值b,　使　|a-b| < radius　　.
    在函数内部打印出所有n个组成员，分成n行输出

    (要求：不能使用数组排序操作)

    例如:
       输入 data[] = { 1, 20, 89, 22, 72, 2,39, 3,56,86, 5, 93,13, 15, 18, 73, 79, 81, 25, 38, 43, 83,48, 52, 59,92,84,95,87 };
       正确的屏幕输出为组及成员为(每行为一组,行之间顺序随意,组内成员顺序随意)：
       1, 2, 3, 5,
       13, 15, 18, 20, 22, 25,
       39, 38, 43, 48, 52, 56, 59,
       73, 72,
       79, 89, 92, 84, 95,87,86, 93,81, 83,

*/
#include <iostream>
#include <vector>
#include <set>

using namespace std;

void cluster(int data[], int len, int radius)
{
    set<int> done; 
    for (int i = 0; i < len; i++)
    {
        if (done.find(data[i]) != done.end()) continue;

        vector<int> g;
        done.insert(data[i]);
        g.push_back(data[i]);

        for (int j = i+1; j < len; ++j) {
            if (done.find(data[j]) != done.end()) continue;

            bool ok = false;
            for (int k = 0; k < g.size(); k++) {
                if (abs(g[k] - data[j]) <= radius) {
                    ok = true;
                    break;
                }
            }
            if (ok) {
                done.insert(data[j]);
                g.push_back(data[j]);
                j = i+1;
            }
        }
        for (vector<int>::iterator it = g.begin(); it != g.end(); it++) {
            cout << *it << ",";
        }
        cout << endl;
    }
}

int main()
{
    // Project 3
    cout << "==========Project3=================" << endl;
    int data[] = { 1, 20, 89, 22, 72, 2,39, 3,56,86, 5, 93,13, 15, 18, 73, 79, 81, 25, 38, 43, 83,48, 52, 59,92,84,95,87 };
    cluster(data, sizeof(data)/sizeof(int), 5);
    /*
    正确的输出为组及成员为(组内成员顺序随意)：
    1, 2, 3, 5,
    13, 15, 18, 20, 22, 25,
    39, 38, 43, 48, 52, 56, 59,
    73, 72,
    79, 89, 92, 84, 95,87,86, 93,81, 83,
    */
    cout << "==========Project3=================" << endl;

    return 0;
}
