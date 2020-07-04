#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <fstream>
#include <sstream>
#include <memory>
#include <stdexcept>

using namespace::std;

string exec(const char* cmd) {
    std::array<char, 128> buffer;
    std::string result;
    std::shared_ptr<FILE> pipe(popen(cmd, "r"), pclose);
    if (!pipe) throw std::runtime_error("popen() failed!");
    while (!feof(pipe.get())) {
        if (fgets(buffer.data(), 128, pipe.get()) != nullptr)
            result += buffer.data();
    }
    return result;
}

void executeZCHAFF(vector<int> &instance, int varCounter, string fileName){
       
    string input = string("glucose_static -model") + " " + fileName;
    string result = exec(input.c_str());

    string prefix("s SATISFIABLE");

    istringstream iss(result);
    stringstream ss;

    string line;
    string token;
    while (getline(iss, line))
    {
            if (!line.compare(0, prefix.size(), prefix)){
                    getline(iss, line);
                    ss << line;
                    ss >> token;
                    instance.push_back(0);
                    for(int i =0; i < varCounter; i++){
                            int variable;
                            ss >> variable;
                            instance.push_back(variable);					
                    }			
                    break;	
            }				
    }
}

vector<string> split(string input){
	vector<string> tokens;
	stringstream ss(input);
	string s;

	while (getline(ss, s, ' ')) {
		tokens.push_back(s);
	}
	return tokens;
}


int fileLineCounter(string fileName){
    int number_of_lines = 0;
    std::string line;
    std::ifstream myfile(fileName);

    while (std::getline(myfile, line)){
    	vector<string> tokens = split(line);
		if (!tokens.empty()) 
			++number_of_lines;	        
    }
        
    myfile.close();
    return number_of_lines;
}


int literalCounter(string fileName, map<string,int> &mapping, vector<string> &mappingBack){
        int variableCounter=0;
        ifstream myfile(fileName);

        string line;
		while (std::getline(myfile, line)){
			vector<string> tokens = split(line);
			for(string token : tokens){
				string literal = (token[0] == '-') ? token.substr(1) : token;
				if(mapping.find(literal) == mapping.end()){
					mapping[literal] = ++variableCounter;
					mappingBack.push_back(literal);
				}
			}
		}
		myfile.close();
		return variableCounter;
}

int main(){
		string line;
		int numberOfLiteral, numberOfClauses;
		vector<string> mappingBack;
		map<string, int> mapping;
		ofstream outputFile;
		
		string inputFileName = "entrada.txt";
		string outputFileName = "formatoCNF.txt";
		
		
		outputFile.open(outputFileName);
		mappingBack.push_back("");
		numberOfClauses = fileLineCounter(inputFileName);
		numberOfLiteral = literalCounter(inputFileName, mapping, mappingBack);
		
        ifstream myfile(inputFileName);
        
        outputFile << "p cnf " << numberOfLiteral << " " << numberOfClauses << endl;
        
		while (std::getline(myfile, line)){
			vector<string> tokens = split(line);
			if (tokens.empty()) continue;
			for(string token : tokens){
				string literal = (token[0] == '-') ? to_string(-1*mapping[token.substr(1)]) + " " : to_string(mapping[token]) + " ";
				outputFile << literal;
			}
			outputFile << 0 << endl;
		}
		
		myfile.close();
		outputFile.close();
		vector<int> instances;
		executeZCHAFF(instances, numberOfLiteral,outputFileName);
		if(instances.size() == 0)
			cout << "UNSAT" << endl;
		for (int e: instances){
			if(e > 0 ){
				cout << mappingBack[e] << endl;
			}
		//	else if (e < 0){
			//	cout << "-" << mappingBack[-e] << endl;
		//	}
		}
		
		
}
