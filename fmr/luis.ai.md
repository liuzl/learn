# luis.ai - Language Understanding Intelligence Service

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
#### 数据标注
#### 特征提取
Token|Token Format|...|Phrase List Matching|Regular Expression Matching
----|----|----|----|----
Schedule|cap||none|none
a|lower||none|none
meeting|lower||none|none
for|lower||none|none
my|lower||none|none
team|lower||none|none
next|lower||none|none
Tuesday|cap||none|none
morning|lower||none|none
with|lower||none|none
Johannes|cap||Name-start|none
at|lower||none|none
13|num||none|Floor-start
F|cap||none|Floor-end

* Gated Recursive Convolutional Neural Networks (grConv)
  * Learning grSemi-CRFs
    * 通过grConv得到不定长phrase的表示
    * 无缝接纳Label transition和Lexicon
    * 计算容易并行
#### 影响LUIS机器学习模型的主要因素
* 标注数据量和质量
* 用户自定义特征 Phrase List & Regex

#### 自定义特征
* Phrase List
  * 通常用于帮助Entity Extraction
  * 提供某一类Entity的可能值，例如：咖啡种类，城市名称等
  * 不需要指定具体对应哪一类Entity，系统会自动推断
  * 须要在标注的数据中有一定的覆盖率（保证被充分训练）
  * 须要能够覆盖未出现在标注数据中的Entity取值（增强通用性）
  * 定义一些Intent相关的关键字可以帮助Intent的识别
* Regular Expressions
  * 和Phrase List类似，只是通过特定规则来定义Entity的取值，例如：时间，比分，订单号等
#### 交互式Phrase List编辑
* 我们从互联网上挖掘出一系列包含特定语义的列表
* 用户只需要提供1~2个词即可开启推荐
* 根据用户对推荐词的选择，系统会不停优化推荐算法
* 实验表明通过交互式编辑，系统可节省50~70%的编辑工作
### 主动学习（Active Learning）
#### 最大化数据标注收益
* 在EndPoint日志中自动选择待标注数据
* 更快的Accuracy增长
#### LUIS中的主动学习
* 推荐对某个Intent/Entity最有帮助的Utterance


## 参考资料
* [AIMOOC 2.0 | 微软智能语义理解服务LUIS技术详解课程](https://mp.weixin.qq.com/s/HwYKhVWPUip6dWwEwYHXIQ)
