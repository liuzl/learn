from collections import defaultdict
import time
import random
import dynet as dy
import numpy as np

w2i = defaultdict(lambda: len(w2i))
t2i = defaultdict(lambda: len(t2i))
UNK = w2i["<unk>"]

def read_dataset(filename):
    with open(filename, "r") as f:
        for line in f:
            tag, words = line.lower().strip().split(" ||| ")
            yield ([w2i[x] for x in words.split(" ")], t2i[tag])

train = list(read_dataset("./data/classes/train.txt"))
w2i = defaultdict(lambda: UNK, w2i)
dev = list(read_dataset("./data/classes/dev.txt"))
nwords = len(w2i)
ntags = len(t2i)

print(nwords, ntags)

model = dy.Model()
trainer = dy.AdamTrainer(model)

W_sm = model.add_lookup_parameters((nwords, ntags))
b_sm = model.add_parameters((ntags))

def calc_scores(words):
    dy.renew_cg()
    b_sm_exp = dy.parameter(b_sm)
    score = dy.esum([dy.lookup(W_sm, x) for x in words])
    return score + b_sm_exp

for ITER in range(100):
    random.shuffle(train)
    train_loss = 0.0
    start = time.time()
    for words, tag in train:
        my_loss = dy.pickneglogsoftmax(calc_scores(words), tag)
        train_loss += my_loss.value()
        my_loss.backward()
        trainer.update()
    print("iter %r: train loss/sent=%.4f, time=%.2fs" % (ITER, train_loss/len(train), time.time()-start))

    test_correct = 0.0
    for words, tag in dev:
        scores = calc_scores(words).npvalue()
        predict = np.argmax(scores)
        if predict == tag:
            test_correct += 1
    print("iter %r: test acc=%.4f" % (ITER, test_correct/len(dev)))
