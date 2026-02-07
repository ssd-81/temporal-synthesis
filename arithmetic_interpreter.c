#include <stddef.h>
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#define INPUT_LENGTH 1000
#define INTEGER 0
#define OPERATOR 1
#define MAX_CAPACITY 1000
typedef struct {
  int type;
  int number_value;
  char operator_value;
} Token;
Token current_token;//global variable with no initialisation
bool is_digit(char c){
  return c>='0'&& c<='9';
}
bool is_operator(char c){
  return c=='+' || c=='-';
}
void parse(int type){
    if(current_token.type!=type){
      printf("Syntax Error: Value passed doesn't follow the grammar");
      exit(-1);
    }
}
void remove_whitespace(char *white_text){
  int count=0;
  for (int i=0;white_text[i];i++){
    if(white_text[i]!=' '){
      white_text[count++]=white_text[i];
    }
  }
   white_text[count]='\0';
}
void merge_digits(char *alone_text){
  size_t length=strlen(alone_text);
  int top=0;
  int i=0;
  char multi_type[]={};
  if (is_digit(alone_text[i])){
    int value=0;
  while (is_digit(alone_text[i]) ) {
    value=value*10+(alone_text[i]-'0');
    i++;
    }  
    top++;
    multi_type[top]=value;
  }else if (is_operator(alone_text[i])) {
    top++;
    multi_type[top]=alone_text[i];
  }else{
    printf("error");
    exit(-1);
  }
}

int identifying_type(char *text,int i){
  if(is_digit(text[i])){
    int value=0;
    while (is_digit(text[i])) {
    value=value*10+(text[i]-'0');
    i++;
    }
    current_token.type=INTEGER;
    current_token.number_value=value;
    return i;
  }
  else if (is_operator(text[i])) {
    current_token.type = OPERATOR;
    current_token.operator_value = text[i];
    return i + 1;
  }
  else {
    printf("Invalid character: %c\n", text[i]);
    exit(-1);
  }
  
}
int interpreter(char *text){
  // int length=sizeof(*text)/sizeof(text[0]);
  remove_whitespace(text);
  int nums[MAX_CAPACITY]={};
  int num_count=0;
  char operators[MAX_CAPACITY]={};
  int op_count=0;
  bool expect_number = true;
  int length=strlen(text);
  int i=0;
  // int identifying_type(char *text, int i);
  while (i<length && text[i]!='\n') {
  i=identifying_type(text,i);
  if(expect_number){
    parse(INTEGER);
    nums[num_count++]=current_token.number_value;
    expect_number=false;
  }else{
    parse(OPERATOR);
    operators[op_count++]=current_token.operator_value;
    expect_number=true;
  }
  }
  int result=nums[0];
  for(int i=0;i<num_count;i++){
    if(operators[i]=='+'){
      result+=nums[i+1];
    }else if(operators[i]=='-'){
      result-=nums[i+1];
    }
  }
  return result;
}
int main(){
  while(true){
  printf(">>>");
  char buf[INPUT_LENGTH];
  fgets(buf,INPUT_LENGTH,stdin);
  int answer=interpreter(buf);
  printf("%d\n",answer);
  };
  
  return 0;
}
