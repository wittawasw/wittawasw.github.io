---
layout: post
title: การเขียน Test ไม่ได้หมายความว่าเราต้องใช้ Test Runner เสมอไป
tags: tdd, testing, python
date: 2024-04-09 13:16 +0700
---
อาทิตย์ก่อนมีโอกาสได้ร่วม training ในหัวข้อ TDD แล้วก็รู้สึกว่าหลายคนมี mindset ประมาณว่า

> ไม่ต้องเขียน test ถ้ายังไม่มี Test Runner

ซึ่งผมมองว่า <u>ผิด</u> โดยความเห็นของผมคือ ต่อให้ใน project ไม่มี test runner
อยู่ เราก็ควรเขียน Unit Test เป็นอย่างน้อยอยู่ดี

หลังจากที่ได้ไป pair กับคนที่ทำ Javascript, Python ที่ไม่เคย setup Test Runner
มาก่อนก็เลยได้ลองเขียน `assertion` แบบง่ายๆ ให้ดูว่าสุดท้ายแล้วการเขียน test ก็แค่
การเขียนฟังก์ชันให้เรียกใช้ฟังก์ชันที่เราต้องการทดสอบและเปรียบเทียบผลลัพธ์เท่านั้นเอง

Code ส่วนของ Javascript ทำที่เครื่องคนอื่นก็เลยไม่ได้เก็บไว้ มีแต่ Python
ที่ทำในเครื่องตัวเองก็เลยยังอยู่ แต่จริงๆแล้วเลือกใช้ ภาษา Dart ตอน training
เพราะเป็นภาษาที่ใช้ในงานปัจจุบัน แต่ยังไม่คล่อง ก็เลยถือโอกาสลองฝึกไปด้วย

```python
from datetime import datetime, timedelta

months = {
  "3": [31, 31],
  "4": [30, 3000],
  "5": [31, 3100],
  "6": [30, 3000],
  "7": [31, 31],
}

months_arr = [31, 3000, 3100, 3000, 31]

def getBudgets(from_date, to_date):
  from_date = datetime.strptime(from_date, "%Y-%m-%d")
  to_date   = datetime.strptime(to_date, "%Y-%m-%d")

  if (to_date - from_date).days < 0:
    return "Error"

  values = []
  current_date = from_date

  while current_date <= to_date:
    values.append(months[str(current_date.month)][1] / months[str(current_date.month)][0])
    current_date += timedelta(days=1)

  return sum(values)

def testGetBudgets(from_date, to_date, expect):
  test_case = "Test: " + from_date + " : " + to_date + " = "

  actual = getBudgets(from_date, to_date)
  if actual == expect:
    return test_case + "Passed"
  else:
    return test_case + "Failed" + "\n" + f"Expected: {expect}, got {actual}"


print(testGetBudgets("2024-07-02", "2024-07-04", 3))
print(testGetBudgets("2024-07-04", "2024-07-02", "Error"))
print(testGetBudgets("2024-06-20", "2024-07-04", 1100 + 4))
print(testGetBudgets("2024-03-10", "2024-07-02", 22 + 9100 + 2))
print(testGetBudgets("2024-06-18", "2024-07-04", 1300 + 4))
```
