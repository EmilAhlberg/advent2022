print([sum(tup) for tup in zip(*[(lambda lb1,lb2,ub1,ub2: (1 if lb1 <= lb2 and ub1 >= ub2 or lb2 <= lb1 and ub2 >= ub1 else 0, 1 if not (lb1 > ub2 or ub1 < lb2) else 0))(int(lb1),int(lb2),int(ub1),int(ub2)) for ([lb1,ub1],[lb2,ub2]) in [(elf1.split('-'), (elf2.split('-'))) for (elf1,elf2) in [row.strip().split(',') for row in open('input/04.txt', 'r')]]])])

# parts that were combined:
#p1_p2_evaluator = lambda lb1,lb2,ub1,ub2: (1 if lb1 <= lb2 and ub1 >= ub2 or lb2 <= lb1 and ub2 >= ub1 else 0, 1 if not (lb1 > ub2 or ub1 < lb2) else 0) 
#boundaries = [(elf1.split('-'), (elf2.split('-'))) for (elf1,elf2) in [row.strip().split(',') for row in open('04.txt', 'r')]]
#result = [sum(tup) for tup in zip(*[evaluator(int(lb1),int(lb2),int(ub1),int(ub2)) for ([lb1,ub1],[lb2,ub2]) in boundaries])
