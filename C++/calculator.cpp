/*******************************************************************************
 *      file name: calculator.cpp                                               
 *         author: Hui Chen. (c) 2019                             
 *           mail: chenhui13@baidu.com                                        
 *   created time: 2019/03/23-16:19:02                              
 *  modified time: 2019/03/23-16:19:02                              
 *******************************************************************************/
#include <iostream>
using namespace std;
/* Project 2
  编写一个函数int Calculate(node *pHead, int k);在链表数字之间插入+ 、- 或者什么都不插入,使得计算结果等于
  给定的整数 K,返回所有可能性的个数。例如链表有数字 1, 2, 3, ..., 9 并且给
  定的整数为 100,那么其中的一种可能性是: 1 + 2 + 34 – 5 + 67 – 8 + 9 = 100。
  (不允许开辟额外的字符串空间,假设链表所有数字都大于 0 并且小于 10)
*/
struct node {
    int val;
    node *next;
    node(int x) : val(x), next(NULL) {}
};

void create_linklist(node **pHead, int n)
{
	*pHead = NULL;
	node *pTail = NULL;
	for (int i=1; i <= n; i++)
	{
		node *p = new node(i);
		if (!(*pHead))
		{
			*pHead = p;
			pTail = p; 
			continue;
		}
		pTail->next = p; 
		pTail = p; 
	}
	pTail->next = NULL;
}
void print_linklist(node *pHead)
{
	node *p = pHead;
	while(p)
	{
		cout<<p<<"	:p->val:"<<p->val<<endl;
		p = p->next;
	}
}
//trace 用户打印函数栈，debug
//op 1 : +
//op -1 : -
//pre用于处理无运算符号，数字直接连接的情况。
//sum: 当前累积的总和。
void cal(node *p, int pre, int sum, int op, int k, int& cnt)
{
	int sum2 = 0; 
	if (op == 1)
		sum2 = sum + pre*10 + p->val; 
	else if (op == -1)
		sum2 = sum - (pre*10 + p->val); 
	if (!p->next)
	{
		//找到符合条件结果
		if (sum2 == k)
			cnt = 1;
		return;
	}
	int c1 = 0, c2 = 0, c3 = 0;
	// case +
	cal(p->next, 0, sum2, 1, k, c1);
	// case -
	cal(p->next, 0, sum2, -1, k, c2);
	// case no operation, contact 
	cal(p->next, pre*10 + p->val, sum, op, k, c3);
	cnt = c1 + c2 + c3;
}
int Calculate(node *pHead, int k)
{
	int count = 0;
	//pre 0, sum 0, op 1 means '+'
	cal(pHead, 0, 0, 1, k, count);
	return count;
}

int main()
{
    // Project 2
    cout << "==========Project2=================" << endl;
	node *pHead;
    //delete_linklist(&pHead);
    create_linklist(&pHead,9);
    cout << Calculate(pHead,100) << endl;
    cout << "==========Project2=================" << endl;
    cout << endl;
}
