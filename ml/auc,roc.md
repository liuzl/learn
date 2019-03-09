# 机器学习模型评估
## 经验误差
* 错误率：`error rate`，分类错误的样本数占样本总数的比例
* 精度：`accuracy`，1-错误率
## 性能度量
### 分类结果混淆矩阵
混淆矩阵 |预测正例|预测反例
------|----|-----
**真实正例**|TP（真正例）|FN（假反例）
**真实反例**|FP（假正例）|TN（真反例）

* 查准率`precision`：$P = \frac{TP} {TP+FP}$
* 查全率`recall`：$R=\frac{TP} {TP+FN}$
* F1：$F1=\frac {2 \times P \times R} {P+R}$
* F1的一般度量形式：$F_{\beta}=\frac {(1+\beta^{2}) \times P \times R} {(\beta^{2} \times P) + R}$
  * 其中$\beta>0$度量了查全率对查准率的相对重要性；
  * $\beta=1$时，退化为标准的F1；
  * $\beta>1$时查全率有更大影响；
  * $\beta<1$时，查准率有更大影响。

### ROC与AUC



## 参考资料
* 周志华，西瓜书第二章 模型评估与选择
* [AUC，ROC我看到的最透彻的讲解](https://blog.csdn.net/u013385925/article/details/80385873)
