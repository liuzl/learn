# luis.ai

## LUIS的技术浅析 - 曹涌

### LUIS的设计理念
* EDI - Enterprise Deep Intelligence，面向企业用户的人工智能助理（Outlook、Skype for bussiness），给EDI发邮件或聊天的形式。
  * 组织会议
  * 设定提醒
  * 园区班车
  * 知识问答

* Meeting Arrangement Example
> Schedule a meeting for my team next Tuesday morning with Johannes at 13F

* LUIS Mission
> Enable non-expert developers to easily create, maintain and integrate high quality language understanding models for their applications.

### 机器学习算法和工作原理
#### LUIS - 机器学习NLU解决方案
* 简单易上手
  * 不需要了解算法细节
  * 内建基本特征
  * 深度学习（Deep Learning）：自动学习高阶语言表示
  * 用户可提供额外特征：Pharse list，正则表达式
* Prebuilt domain/entity
  * 基本工具：时间，词性等标注器
  * 基本领域：Calendar, Events, Music etc.
* 最大化数据标注收益
  * 主动学习（Active Learning）

#### 定义Intent和Entity
* Intent
  * Understanding User Intent
  * 我的应用可以为用户做什么？
  * Function
  * Text Classification
* Entity
  * Detecting Task-related Entities
  * 我们需要用户提供哪些参数？
  * Parameters
  * Entity Extraction
### 主动学习


## 参考资料
* [AIMOOC 2.0 | 微软智能语义理解服务LUIS技术详解课程](https://mp.weixin.qq.com/s/HwYKhVWPUip6dWwEwYHXIQ)
