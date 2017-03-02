// https://leetcode.com/problems/lru-cache/?tab=Description
#include <iostream>
#include <unordered_map>
#include <list>

using namespace std;

class LRUCache {
public:
    LRUCache(int capacity) {
        capacity_ = capacity;
        data_.clear();
        used_.clear();
    }
    
    int get(int key) {
        if (data_.find(key) == data_.end()) {
            return -1;
        }
        used_.remove(key);
        used_.push_back(key);
        return data_[key];
    }
    
    void put(int key, int value) {
        if (data_.find(key) != data_.end()) {
            data_[key] = value;
            used_.remove(key);
            used_.push_back(key);
        } else {
            if (data_.size() >= capacity_) {
                data_.erase(used_.front());
                used_.pop_front();
            }
            data_[key] = value;
            used_.push_back(key);
        }
    }
private:
    unordered_map<int, int> data_;
    list<int> used_;
    int capacity_;
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */

int main(int argc, char* argv[])
{
    LRUCache cache( 2 /* capacity */ );

    cache.put(1, 1);
    cache.put(2, 2);
    cout << cache.get(1) << endl;       // returns 1
    cache.put(3, 3);    // evicts key 2
    cout << cache.get(2) << endl;       // returns -1 (not found)
    cache.put(4, 4);    // evicts key 1
    cout << cache.get(1) << endl;       // returns -1 (not found)
    cout << cache.get(3) << endl;       // returns 3
    cout << cache.get(4) << endl;       // returns 4
    return 0;
}
