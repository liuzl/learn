#include <iostream>

using namespace std;

int main(int argc, char* argv[]) {
    auto basicLambda = [] { cout << "Hello, world!" << endl; };
    basicLambda();
}
