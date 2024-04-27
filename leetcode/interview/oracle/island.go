package oracle

func findIsland(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	rowLen, colLen := len(matrix), len(matrix[0])

	nIsland := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				nIsland++

				q := Queue{
					arr: make([][2]int, 0),
				}
				q.Enqueue([2]int{i, j})

				for q.Size() > 0 {
					element := q.Dequeue()
					row, col := element[0], element[1]

					matrix[row][col] = -nIsland

					distance := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

					for k := 0; k < len(distance); k++ {
						new_row, new_col := row+distance[k][0], col+distance[k][1]

						if new_row >= 0 && new_row < rowLen && new_col >= 0 && new_col < colLen && matrix[new_row][new_col] == 1 {
							q.Enqueue([2]int{new_row, new_col})
						}
					}
				}
			}
		}
	}

	return nIsland
}

type Queue struct {
	arr [][2]int
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) Enqueue(x [2]int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Dequeue() [2]int {
	first := [2]int{}
	if q.Size() > 0 {
		first = q.arr[0]
	}
	if q.Size() > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = [][2]int{}
	}

	return first
}

/*

interface ReportGenerationService{
     generatePdfReport();
     generateXmlReport();
     generateCsvReport();
}

class ReportGenerationServiceImp implements ReportGenerationService{
    Report  generatePdfReport(){--}
    Report  generateXmlReport(){--}
    Report  generateCsvReport(){--}
}

class ReportGenerationServiceImpAdapter implements ReportGenerationService{
	ReportGenerationNew serviceNew = new ReportGenerationServiceImplNew();

    Report  generatePdfReport(){
		serviceNew.generatePdfReportNew()
	}

    Report  generateXmlReport(){
		serviceNew.generateXmlReportNew()
	}
    Report  generateCsvReport(){--}
}

public class ReportGenerator{
Private ReportGenerationService service;

ReportGenerator(ReportGenerationService service){
this .service=service;
}

Public Report getPdfReport(){
   return service.generatePdfReport();
}

Public Report getXmlReport(){
   return service.generateXmlReport();
}

Public Report getCsvReport(){
   return service.generateCsvReport();
}


}

Public class ReportGeneratorMain{
 public static void main(String[] args){
//ReportGenerationService service = new ReportGenerationServiceImpl();
ReportGenerationService service = new ReportGenerationServiceImplAdpater();
ReportGenerator repGen = new ReportGenerator(service);
Report report = repGen.getPdfReport();
----
----

}
}

SOLID -> Interface Seggregation, Dependency Injection

ACID -> Isolation ->
*/
