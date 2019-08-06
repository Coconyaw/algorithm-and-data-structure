# AOJ ALDS1_6_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_6_C

class Card():
    def __init__(self, suit: chr, number: int):
        self.suit = suit
        self.number = number

def partition(a: list, left: int, right: int) :
    std = a[right].number
    j = left - 1
    for i, v in enumerate(a[left:right]):
        idx = i + left
        if a[idx].number <= std:
            j += 1
            a[j], a[idx] = a[idx], a[j]

    j += 1
    a[j], a[right] = a[right], a[j]
    return j

def quick_sort(a: list, left: int, right: int):
    if left < right:
        sep = partition(a, left, right)
        quick_sort(a, left, sep-1)
        quick_sort(a, sep+1, right)

def merge(a: list, left: int, mid: int, right: int):
    n1 = mid - left
    n2 = right - mid

    L = []
    R = []
    for i in range(n1):
        L.append(a[left+i])
    for i in range(n2):
        R.append(a[mid+i])

    inf = 1 << 32
    L.append(Card("Z", inf))
    R.append(Card("Z", inf))

    i = 0
    j = 0
    for k in range(left, right, 1):
        if L[i].number <= R[j].number:
            a[k] = L[i]
            i += 1
        else:
            a[k] = R[j]
            j += 1

def merge_sort(a: list, left: int, right: int):
    if left + 1 < right:
        mid = (left + right) // 2
        merge_sort(a, left, mid)
        merge_sort(a, mid, right)
        merge(a, left, mid, right)

def is_stable(m, q):
    for i, j in zip(m, q):
        if i.number != j.number or i.suit != j.suit:
            return False
    return True

n = int(input())

cards = []
for i in range(n):
    card_info = input().split(" ")
    cards.append(Card(card_info[0], int(card_info[1])))

merge_sorted = cards.copy()
quick_sorted = cards.copy()

merge_sort(merge_sorted, 0, len(merge_sorted))
quick_sort(quick_sorted, 0, len(quick_sorted) - 1)
stable = is_stable(merge_sorted, quick_sorted)

if stable:
    print("Stable")
else:
    print("Not stable")

for c in quick_sorted:
    print("{} {}".format(c.suit, c.number))

# print("##############")
# for c in merge_sorted:
#     print("{} {}".format(c.suit, c.number))
