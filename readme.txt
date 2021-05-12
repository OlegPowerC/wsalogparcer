wsalogparcer.exe -s <start time> -e <end time> -u <username or part of the name> -i <IP address> -a <URL> -f <log file>

Time format is: DD/MM/YYYY/HH/MinMin/SS/TimeZone
Example 23/02/2017/12:32:00/MSK

Username: can be full or part like: alex al or Ale

IP address, full or part: example: 10.111.18.

URL: can be full or part: example: google

Mandatory flag -f log file name (ACL log from Cisco WSA)

Example:
wsalogparcer.exe -s 01/01/2018/00:00:00/MSK -e 02/07/2018/13:49:59/MSK -u Yulia  -f test.txt

Output:
Date and time
IP address
Request type (GET/POST e.t.c.)
URL
Username
Policy and more detail

Example:
2018-10-12 17:20:11.000000707 +0300 MSK
10.111.18.166
GET
http://blob.weather.microsoft.com/static/weather4/ru/txw/26.png
"TESTAD\yulia.ivanova@TESTAD"
ALLOW_CUSTOMCAT_12-TESTADADPol-TESTADADUSERS-NONE-NONE-NONE-DefaultGroup
ALLOW_CUSTOMCAT_12
TESTADADPol
TESTADADUSERS

--------------------------------------------------------------------------------------------------------------------------------------------

wsalogparcer.exe -s <����� ������ ������> -e <����� �� �������� ������> -u <��� ������������> -i <IP �����> -a <����� �� URL> -f <��� ����>

����� � ������� ����/�����/���/����/������/�������/����
������ 23/02/2017/12:32:00/MSK

��� ������������: ����� ��������� �����, �������� alex

IP �����, ����� ��������� �����: �������� 10.111.18.

����� �� URL: ����� ������� ����� ������: �������� google

������������ �������� -f ��� ����� (ACL ��� � Cisco WSA)

������:
wsalogparcer.exe -s 01/01/2018/00:00:00/MSK -e 02/07/2018/13:49:59/MSK -u Yulia  -f test.txt


�����:
���� � �����
IP �����
��� ������� (GET/POST � �.�.)
URL
��� ������������
�������� � ������

������:
2018-10-12 17:20:11.000000707 +0300 MSK
10.111.18.166
GET
http://blob.weather.microsoft.com/static/weather4/ru/txw/26.png
"TESTAD\yulia.ivanova@TESTAD"
ALLOW_CUSTOMCAT_12-TESTADADPol-TESTADADUSERS-NONE-NONE-NONE-DefaultGroup
ALLOW_CUSTOMCAT_12
TESTADADPol
TESTADADUSERS