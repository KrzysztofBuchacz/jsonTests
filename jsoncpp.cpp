#include <iostream>
#include <fstream>
#include <json/json.h>

int main()
{
    std::cout << "JOO!" << std::endl;


    try
    {
        Json::Value obj;
        obj["jojo"] = std::numeric_limits<double>::infinity();
        std::ofstream fileStream("badInfCPP.json");
        Json::StreamWriterBuilder writer;
        fileStream << Json::writeString(writer, obj);

        std::ifstream file("../badInf.json");
        Json::CharReaderBuilder reader;
        std::string errs;
        std::cout << Json::parseFromStream(reader, file, &obj, &errs) << std::endl;
        std::cout << "Errs: " << errs << std::endl;
    }
    catch (...)
    {
        std::cout << "Ku!";
    }

    return 0;
}
