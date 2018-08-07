# Notes on "Building new NLP solutions with spaCy and Prodigy"

## Essential Questions from the Machine Learning Hierarchy of Needs
1. How will the model fit into the larger application or business process?
   1. What is the *decision* that will be impacted by a machine learning solution?
   2. Is there an existing software application we can plug into? Can we develop an API?
   3. Do we need to build something entirely custom?
2. What is the annotation scheme? What is being labeled? How will the corpus be built?
   1. Are we classifying documents?
   2. Are we tagging spans of text?
   3. Are we trying to automatically generate structured data from text?
      * What models do we need to build and compose to generate this data?
3. What quality control processes will we use to ensure consistent and clean data?
    1. Who are experts who we can get to label data accurately?
    2. How will we construct a gold standard data set for evaluation?
    3. How can we break down the NLP problem to improve the ease and consistency of the annotation task?

## Things to Avoid
1. **Overambitious *Imagineering*** - Thinking too big, rather than specifically about what business problem(s) can be solved.
2. **Forecasting without experimentation** - Most problems will need to be evaluated after experimentation with data: don’t put excessive value on arbitrary thresholds of accuracy or other performance metrics up front
3. **Outsourcing** - You want a dataset and annotation process that is reliable and consistent, not noisy. Outsourcing solves the wrong problem (marginal annotation cost).
4. **Overengineering & AI FOMO** - Can you compose existing, generic models to solve your task rather than tweaking parameters on the latest algorithms?
5. **Premature Shipping** - if it fails in development, it certainly won’t work in production.

## General Principles (with timestamp)
* Iterate on your data (labeling, annotation scheme, problem framing), as well as your code. (10:52)
* Rather than make assumptions, get to iterative experimentation as fast as possible. (11:14)
* Start with what you need at output, then work towards breaking down the problem into a series of modeling decisions. (12:17)
* Break down the problem so that your component models can perform accurately and you can cheaply generate annotations. (14:00) Compose generic models into novel solutions (17:00)
  * *Example*: Is this document about a crime? If so, what do the entities represent? By comparison: jumping straight to classifying the victim, perpetrator, location is both hard for a model to classify and cognitively burdensome for a human to annotate.
* Fine-tuning existing models for your domain, using generic categories, is much cheaper than starting from scratch (17:20)
  * *Example*: fine tuning the dependency parsing on conversational or colloquial text might improve the accuracy. (33:50)
* Annotate events and topics at the sentence, paragraph, or document level to avoid boundary issues. Additionally, annotation is easier and classification is more accurate. (17:50)
* Bootstrap your project estimates based on contact with the evidence, rather than making assumptions or guesses up front. (23:17)
* Reframe your problem so that the annotator has to consider less information at once: use pre-suggested outputs and make corrections. This reduces error by asking the annotator to correct the model when its wrong (vs. getting tired/lazy on annotation tasks from scratch) (28:57)

## References
* Honnibal, M. (2018, July). *[Building new NLP solutions with spaCy and Prodigy](https://speakerdeck.com/honnibal/building-new-nlp-solutions-with-spacy-and-prodigy)*, Presented at PyData Berlin 2018, Berlin, Germany.
* https://pmbaumgartner.github.io/blog/notes-on-nlp-projects/
