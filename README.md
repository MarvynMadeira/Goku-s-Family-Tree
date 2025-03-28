#Step by Step

- first:
    - I create a familyMember struck with elements: 
    - name string 
    - wife *familyMember (of member)
    - children []*familyMember (of member too)
    - father   *familyMember 
    - mother   *familyMember
    
    - The wife's slice is a nill because some members of family isn't married. The same is applies to children.
    - [] in the children it is possible that the number of children is not fixed.
    - Why not an array instead of a slice? Arrays in go have fixed values when defined in a declaration. Slices have the possibility of manageable and varied values.

- Second:
    - I create func newMember. 
    - Receives name, father and mother, returns it within the 'familyMember' struct adding a new member in the tree.
    - & in '&familyMember' it is used to point to the memory address of a var, it creates a pointer to the var. 

- Third:
 - I create func (s *familyMember) newChildren(newPerson *familyMember).
 - s (s - struct, i who chose to) is a receiver of method.
 - (s *familyMember) is a receiver of method, the familymember's instance which summon the method (newChildren).
 - append is used to add newPerson to Children.
 - newPerson.father = s points to the father.

- Fourth:
  - I create func (s *familyMember) show_me_thisFamily(tree int).
  - Tree is an int type param, which indicates the depth level of the struct.
  - for i := 0; i < tree; i++:
    - for is a beginning of a loop.
    - i := 0 is var by value = 0.
    - The loop will continue as long as the value of 'i' is less than tree (i < tree).
    - i++ with each interaction add value 1 more.
    - for _, indicates that only values will be exposed.
    - range indicates that it will show the collection, in this case, children.
    - tree + 1 increases one level with each new member.

- Fifth:
  - The output should be: 
    - Goku's Family Tree:
          - Bardock
            - Goku
                - Gohan 
                    - Pan
                - Goten
            - Raditz
  
