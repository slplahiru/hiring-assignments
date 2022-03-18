# Assignment used for Data Scientist Position

This task is intended for candidates applying for a Data Science position at Visma. The assignment contains real data and directly reflects the actual challenges we face at Visma when we are trying to apply machine learning to the field of automating accounting processes.

Your task is to (attempt to) improve on a baseline included with this repository - using a slightly different approach.

## The problem
Companies often have a lot of expenses. Each payment of these expenses is in accounting terms called a financial transaction which should be entered into the financial books. But where? This is the job of the accountant to figure out. An expense for paint might go to the maintenance account, and an expense for a taxi ride might go to the account for travel expenses. Each company will have a set of accounts they use to categorize their spending - the _chart of accounts_.

In Visma’s systems it is possible to import your bank statements and use these to create the financial transactions, so all you have to do is to import the bank statements, and decide on which account you want to book each expense ..and you’re done.

But what if we could make this process even smoother by learning how the expenses should be booked? Wouldn’t that be great? If we know that a company's books looks like this:


    “Cleaning service & Co 45JW5PP12”          --> Maintenance account
    “Andersen’s cleaning service 02.02.2017”   --> Maintenance account
    “McDonald's Nørrebro”                      --> Employee catering account

We might just go ahead and help the company do future bookings of _“McDonald's Nørrebro”_, or maybe we can assist the company with new like _“Copenhagen Cleaning Company”_ if we build a clever model, or maybe we might even be able to help with _“Burger King Kastrup Airport”_ if we look at how other companies usually do their bookkeeping.

Here are some of the important characteristics of this problem:

* The same text from the bank statements might lead to two different accounts. An example of this is that you might buy both food for your employees (which you might want to put on the Employee catering account) and chocolate (which you might want to put on the Gifts and flowers account) in the same shop.
While each company has a unique chart of accounts - a lot of companies use very similar charts
* However many companies have a chart of accounts that varies only a little from the default chart of accounts.
* Two companies that have the same chart of accounts might not agree how certain expenses should be booked. An example of this might be sandwiches bought for a meeting with business partners. One company might always put this on the _Employee catering_ account, another one might always put it on the _Meetings_ account, and a third might even have a _Sandwiches for meetings with business partners_ account.
* Each account has a number. There are no restrictions on these account numbers besides that an account number is unique within a company. So the Travel expense account in one company might have the same number as the Employee catering account in another company.

## Part 1 - Get the baseline up and running

Included with this repository is a basic scikit-learn model for this problem. You find it in `model.py` - train the model on the data included and keep it as a baseline.

## Part 2 - Improve with a modern model

Now you have a baseline - but perhaps we can take a little further with new techniques. Using Keras (or similar - see below)*, create a simple neural network for the same problem. Some ideas you might want to consider
* Can you improve on the text treatment with something like a word embedding?
* Is a little nonlinearity helpful?
* Is CompanyId useful?

Some guidelines: 

* Don’t overthink it - this is a simple dataset - and all text is obfuscated so fancy modern NLP with pretrained foundation models will not help you. 
* What we’re looking for is a neat and tidy homegrown model with clean code.
* Make use of any Keras standard layer or function you find useful. Don’t reinvent any wheel Keras already built for you. 

*If you don’t really like Keras - but are familiar with another framework - Pytorch or JAX or similar - feel free to use that instead.

## Part 3 - Compare and discuss

Evaluate your results - compare them against the baseline. What are your conclusions? Any ideas for next steps?


## Description of data
The dataset consist of expenses from 100 random companies. Customers of one of Visma's ERPs. For each company we provide all expenses that was booked in the ERP.

Description of each column in the dataset:
- __CompanyId:__ The identifier of the company to help you slice and dice the data in the right way.
- __BankEntryDate *(feature)*__: The date of the financial transaction.
- __BankEntryText *(feature)*__: The text following along with the financial transaction. This is typically machine generated, but in case of manual transactions they may be manually written by a human. _Please note that the text has been split into words before they have been hashed._
- __BankEntryAmount *(feature)*__: The amount of the financial transaction. Expenses are negative, earnings are positive.
- __AccountNumber *(target)*__: The account number. The uniquely identifies an account, and can therefore be used as the target variable / the class that we want to predict.
- __AccountName__: The name of the account.
- __AccountTypeName__: The type of the account.

Columns marked by _(feature)_ can optionally be used as a feature in your predictive model. All of these features are typically what you see when you look at your bank statement.

The __CompanyId__ is special - it basically slices the CSV into 100 separate smaller datasets. Feel free to experiment with solving for all companies with a single model.

The _AccountNumber_ is your target variable. The _AccountName_ and the _AccountTypeName_ are properties of the account, and hence not of direct interest to the prediction problem, but if you can come up with creative ways of using it, then feel free to do so.

The rows are sorted first by _BankEntryDate_, then by _CompanyId_.

For obvious privacy reasons the amounts has been bucketed and the texts has been obfuscated using the following function:

    data = query(limit = 100) # Pandas DataFrame

    def short_hash(word):
        try:
            int(word)
            typ = 'int'
        except:
            typ = 'str'
        bytes_ = word.encode() + secret_salt
        sha_word = hashlib.sha512(bytes_).hexdigest()
        return '{}:{}'.format(typ, sha_word[:7])

    def obfuscate_text(string_):
        return " ".join([short_hash(w) for w in string_.split()])

    def modify_row(row):
        # Translate AccountTypeName to english
        row['AccountTypeName'] = 'Profit and Loss' \
            if row['AccountTypeName'] == 'Drift' else 'Balance'
        # Obfuscate AccountName
        row['AccountName'] = short_hash(row['AccountName'])
        # Obfuscate BankEntryText
        row['BankEntryText'] = obfuscate_text(row['BankEntryText'])
        # Obfuscate CompanyId
        row['CompanyId'] = short_hash(row['CompanyId'])

        p_bar.update()
        return row

    data = data.apply(modify_row, axis=1)

    # Bin BankEntryAmount
    data['BankEntryAmount'] = pd.cut(
        data['BankEntryAmount'],
        bins=[float('-inf'), -10000, -1000, -100, -10, 0, 10, 100, 1000,
              10000, float('inf')],
        labels=['big negative', '> -10000', '> -1000', '> -100', '> -10',
                '< 10', '< 100', '< 1000', '< 10000', 'big positive']
    )

    data.to_csv(output_filename)

The data is a zipped `.csv` file called `bank_expenses_obfuscated.csv.zip`.

## Data example

Here's the top three rows from the data set:

|   | CompanyId   | BankEntryDate | BankEntryText           | BankEntryAmount | AccountName | AccountNumber | AccountTypeName |
|---|-------------|---------------|-------------------------|-----------------|-------------|---------------|-----------------|
| 0 | int:a055470 | 2016-02-29    | str:6cd08e4 int:49fed34 | > -1000         | str:1e82557 | 9900          | Balance         |
| 1 | int:a055470 | 2016-02-29    | str:6cd08e4 int:49fed34 | > -1000         | str:9ce853c | 3115          | Profit and Loss |
| 2 | int:a055470 | 2016-02-29    | str:38248d2             | > -100          | str:a9f0788 | 2240          | Profit and Loss |

## Guidelines
We would like to remind you of a few important things:
- **Important:** This data it real. An acceptable model might not exist, so don't feel bad if your results are disappointing.
- Focus on the right stuff. Don't spend many hours on data wrangling and other stuff that does not show us your true skill-set. Instead, please make a few assumptions, and make sure to tell us about the assumptions you made.
- We do not judge you on the accuracy of your predictive model, but on your problem solving skills. So don't spend all your time tweeking parameters.
- If you feel that you want to know more about the usecase so the you can better derive the external requirements (like the maximum response time at prediction time, or the importance of model interpretability)? Then you can either make up your own requirements assumptions (remember to tell us about these), or ask us.
- Use whatever tech stack you feel comfortable using.

## Got stuck?
You can always email us and ask for advice or just ask question to ensure you correctly understood the task. This will not be seen as a sign of weakness, to the contrary it shows that fully understanding the problem is important to you.

## Suggestions for improvements?

Please help us improve this assignment by suggesting changes, or making a pull request.
