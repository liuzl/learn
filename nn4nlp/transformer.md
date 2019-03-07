# Transformer

## Attention

An attention model differs from a classic sequence-to-sequence model in two main ways:

* The encoder passes a lot more data to the decoder. Instead of passing the last hidden state of the encoding stage, the encoder passes all the hidden states to the decoder.

* An attention decoder does an extra step before producing its output. In order to focus on the parts of the input that are relevant to this decoding time step, the decoder does the following:
  1. Look at the set of encoder hidden states it received – each encoder hidden states is most associated with a certain word in the input sentence.
  2. Give each hidden states a score (let’s ignore how the scoring is done for now).
  3. Multiply each hidden states by its softmaxed score, thus amplifying hidden states with high scores, and drowning out hidden states with low scores.

## 参考资料

* [图解什么是 Transformer](https://www.jianshu.com/p/e7d8caa13b21)
